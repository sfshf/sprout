package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sfshf/sprout/app/govern/bll"
	"github.com/sfshf/sprout/app/govern/config"
	"github.com/sfshf/sprout/app/govern/ginx"
	"github.com/sfshf/sprout/app/govern/schema"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Staff) GetPicCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	resp, err := a.bll.GeneratePicCaptchaIdAndB64s(ctx, c.Query("obsolete_id"))
	if err != nil {
		ginx.JSONWithFailure(c, nil)
		return
	}
	ginx.JSONWithSuccess(c, resp)
	return
}

func (a *Staff) GetPicCaptchaAnswer(c *gin.Context) {
	ctx := c.Request.Context()
	sessionId := ginx.SessionIdFromGinX(c)
	if sessionId.Hex() != config.C.Root.SessionId {
		ginx.JSONWithUnauthorized(c, schema.ErrUnauthorized.Error())
		return
	}
	answer := a.bll.GetPicCaptchaAnswer(ctx, c.Param("id"))
	if answer == "" {
		ginx.JSONWithInvalidCaptcha(c, schema.ErrInvalidCaptcha.Error())
		return
	}
	ginx.JSONWithSuccess(c, answer)
	return
}

func (a *Staff) SignIn(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.SigninReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if config.C.PicCaptcha.Enable && !a.bll.VerifyPictureCaptcha(ctx, arg.PicCaptchaId, arg.PicCaptchaAnswer) {
		ginx.JSONWithInvalidCaptcha(c, schema.ErrInvalidCaptcha.Error())
		return
	}
	staff, err := a.bll.VerifyAccountAndPassword(ctx, arg.Account, arg.Password)
	if err != nil {
		ginx.JSONWithInvalidAccountOrPassword(c, nil)
		return
	}
	clientIp := model.StringPtr(c.ClientIP())
	if staff.SignInIpWhitelist != nil {
		var validIp bool
		for _, ip := range staff.SignInIpWhitelist {
			if ip == *clientIp {
				validIp = true
				break
			}
		}
		if !validIp {
			ginx.JSONWithUnauthorized(c, schema.ErrUnauthorized.Error())
			return
		}
	}
	signInTime := primitive.DateTime(arg.Timestamp)
	resp, err := a.bll.SignIn(ctx, staff.ID, clientIp, &signInTime)
	if err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, resp)
	return
}

func (a *Staff) SignOut(c *gin.Context) {
	ctx := c.Request.Context()
	sessionId := ginx.SessionIdFromGinX(c)
	if err := a.bll.SignOut(ctx, sessionId); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Staff) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.SignupReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.SignUp(ctx, &arg); err != nil {
		ginx.JSONWithDuplicateEntity(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Staff) SignOff(c *gin.Context) {
	ctx := c.Request.Context()
	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	sessionId := ginx.SessionIdFromGinX(c)
	if (objId.Hex() != sessionId.Hex() && sessionId.Hex() != config.C.Root.SessionId) ||
		(sessionId.Hex() == config.C.Root.SessionId && objId.Hex() == config.C.Root.SessionId) {
		ginx.JSONWithUnauthorized(c, schema.ErrUnauthorized.Error())
		return
	}
	if err := a.bll.SignOff(ctx, &objId); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Staff) Update(c *gin.Context) {
	ctx := c.Request.Context()
	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	sessionId := ginx.SessionIdFromGinX(c)
	if objId.Hex() != sessionId.Hex() || sessionId.Hex() != config.C.Root.SessionId {
		ginx.JSONWithUnauthorized(c, schema.ErrUnauthorized.Error())
		return
	}
	var arg bll.StaffUpdateReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.Update(ctx, sessionId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Staff) Profile(c *gin.Context) {
	ctx := c.Request.Context()
	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	res, err := a.bll.Profile(ctx, &objId)
	if err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, res)
	return
}

func (a *Staff) List(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.StaffListReq
	if err := c.ShouldBindQuery(&arg); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	sort := make(bson.M, 0)
	if arg.OrderBy != nil {
		orderBy, err := arg.OrderBy.Values()
		if err != nil {
			ginx.JSONWithInvalidArguments(c, schema.ErrInvalidArguments.Error())
			return
		}
		for k, v := range orderBy {
			sort[k] = v
		}
	}
	res, err := a.bll.List(ctx, &arg, sort)
	if err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, res)
	return
}
