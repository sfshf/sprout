package api

import (
	"github.com/gin-gonic/gin"
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
	resp, err := a.bll.GeneratePicCaptchaIdAndB64s(ctx, c.Query("id"))
	if err != nil {
		ginx.AbortWithFailure(c, nil)
		return
	}
	ginx.JSONWithSuccess(c, resp)
	return
}

func (a *Staff) GetPicCaptchaAnswer(c *gin.Context) {
	ctx := c.Request.Context()
	sessionId := ginx.SessionIdFromGinX(c)
	if sessionId.Hex() != config.C.Root.SessionId {
		ginx.AbortWithUnauthorized(c, schema.ErrUnauthorized.Error())
		return
	}
	answer := a.bll.GetPicCaptchaAnswer(ctx, c.Param("id"))
	if answer == "" {
		ginx.AbortWithInvalidCaptcha(c, schema.ErrInvalidCaptcha.Error())
		return
	}
	ginx.JSONWithSuccess(c, answer)
	return
}

func (a *Staff) SignIn(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.SigninReq
	if err := c.ShouldBind(&arg); err != nil {
		ginx.AbortWithInvalidArguments(c, err.Error())
		return
	}
	if config.C.PicCaptcha.Enable && !a.bll.VerifyPictureCaptcha(ctx, arg.PicCaptchaId, arg.PicCaptchaAnswer) {
		ginx.AbortWithInvalidCaptcha(c, schema.ErrInvalidCaptcha.Error())
		return
	}
	staff, err := a.bll.VerifyAccountAndPassword(ctx, arg.Account, arg.Password)
	if err != nil {
		ginx.AbortWithInvalidAccountOrPassword(c, nil)
		return
	}
	clientIp := model.StringPtr(ginx.GetClientRealIp(c))
	if staff.SignInIpWhitelist != nil {
		var validIp bool
		for _, ip := range staff.SignInIpWhitelist {
			if ip == *clientIp {
				validIp = true
				break
			}
		}
		if !validIp {
			ginx.AbortWithUnauthorized(c, schema.ErrUnauthorized.Error())
			return
		}
	}
	signinTime := primitive.DateTime(arg.Timestamp)
	resp, err := a.bll.SignIn(ctx, staff.ID, clientIp, &signinTime)
	if err != nil {
		ginx.AbortWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, resp)
	return
}

func (a *Staff) SignOut(c *gin.Context) {
	ctx := c.Request.Context()
	sessionId := ginx.SessionIdFromGinX(c)
	if err := a.bll.SignOut(ctx, sessionId); err != nil {
		ginx.AbortWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Staff) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.SignupReq
	if err := c.ShouldBind(&arg); err != nil {
		ginx.AbortWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.SignUp(ctx, &arg); err != nil {
		ginx.AbortWithDuplicateEntity(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Staff) SignOff(c *gin.Context) {
	ctx := c.Request.Context()
	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.AbortWithInvalidArguments(c, err.Error())
		return
	}
	sessionId := ginx.SessionIdFromGinX(c)
	if (objId.Hex() != sessionId.Hex() && sessionId.Hex() != config.C.Root.SessionId) ||
		(sessionId.Hex() == config.C.Root.SessionId && objId.Hex() == config.C.Root.SessionId) {
		ginx.AbortWithUnauthorized(c, schema.ErrUnauthorized.Error())
		return
	}
	if err := a.bll.SignOff(ctx, &objId); err != nil {
		ginx.AbortWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Staff) Update(c *gin.Context) {
	ctx := c.Request.Context()
	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.AbortWithInvalidArguments(c, err.Error())
		return
	}
	sessionId := ginx.SessionIdFromGinX(c)
	if objId.Hex() != sessionId.Hex() || sessionId.Hex() != config.C.Root.SessionId {
		ginx.AbortWithUnauthorized(c, schema.ErrUnauthorized.Error())
		return
	}
	var arg bll.StaffUpdateReq
	if err := c.ShouldBind(&arg); err != nil {
		ginx.AbortWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.Update(ctx, sessionId, &arg); err != nil {
		ginx.AbortWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Staff) Profile(c *gin.Context) {
	ctx := c.Request.Context()
	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.AbortWithInvalidArguments(c, err.Error())
		return
	}
	res, err := a.bll.Profile(ctx, &objId)
	if err != nil {
		ginx.AbortWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, res)
	return
}

func (a *Staff) List(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.StaffListReq
	if err := c.ShouldBindQuery(&arg); err != nil {
		ginx.AbortWithInvalidArguments(c, err.Error())
		return
	}
	sort := make(bson.M, 0)
	if arg.OrderBy != nil {
		orderBy, err := arg.OrderBy.Values()
		if err != nil {
			ginx.AbortWithInvalidArguments(c, schema.ErrInvalidArguments.Error())
			return
		}
		for k, v := range orderBy {
			sort[k] = v
		}
	}
	res, err := a.bll.List(ctx, &arg, sort)
	if err != nil {
		ginx.AbortWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, res)
	return
}
