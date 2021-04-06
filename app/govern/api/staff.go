package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/bll"
	"github.com/sfshf/sprout/app/govern/conf"
	"github.com/sfshf/sprout/gin/ginx"
	"github.com/sfshf/sprout/gin/middleware"
	"github.com/sfshf/sprout/model"
	"github.com/sfshf/sprout/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Staff) GetPicCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	captchaId := c.Query("id")
	resp, err := a.bll.GeneratePicCaptchaIdAndB64s(ctx, captchaId)
	if err != nil {
		ginx.AbortWithFailure(c, nil)
		return
	}
	ginx.JSONWithSuccess(c, resp)
	return
}

func (a *Staff) GetPicCaptchaAnswer(c *gin.Context) {
	ctx := c.Request.Context()
	sessionId := middleware.SessionIdFromGinX(c)
	if sessionId != conf.C.Root.SessionId {
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
	if conf.C.PicCaptcha.Enable && !a.bll.VerifyPictureCaptcha(ctx, arg.PicCaptchaId, arg.PicCaptchaAnswer) {
		ginx.AbortWithInvalidCaptcha(c, schema.ErrInvalidCaptcha.Error())
		return
	}
	staff, err := a.bll.VerifyAccountAndPassword(ctx, arg.Account, arg.Password)
	if err != nil {
		ginx.AbortWithInvalidAccountOrPassword(c, nil)
		return
	}
	clientIp := model.StringPtr(ginx.GetClientRealIp(c))
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
	//ctx := c.Request.Context()
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
	id := c.Param("id")
	staffId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ginx.AbortWithInvalidArguments(c, err.Error())
		return
	}
	sessionId := middleware.SessionIdFromGinX(c)
	if (id != sessionId && sessionId != conf.C.Root.SessionId) ||
		(sessionId == conf.C.Root.SessionId && id == conf.C.Root.SessionId) {
		ginx.AbortWithUnauthorized(c, schema.ErrUnauthorized.Error())
		return
	}
	if err := a.bll.SignOff(c, staffId); err != nil {
		ginx.AbortWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Staff) Update(c *gin.Context) {

}

func (a *Staff) List(c *gin.Context) {

}

func (a *Staff) Info(c *gin.Context) {

}
