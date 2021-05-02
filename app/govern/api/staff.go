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

// GetPicCaptcha
// @description Get a new picture captcha id and base64 string of the picture, and delete the obsolete captcha of the obsolete_id, if has.
// @id get-pic-captcha
// @tags staff
// @summary Get a picture captcha.
// @produce json
// @param obsolete_id query string false "an obsolete captcha id."
// @success 2000 {object} bll.GetPicCaptchaResp "captcha id and picture."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /picCaptcha [GET]
func (a *Staff) GetPicCaptcha(c *gin.Context) {
	ctx := c.Request.Context()
	resp, err := a.bll.GeneratePicCaptchaIdAndB64s(ctx, c.Query("obsolete_id"))
	if err != nil {
		ginx.JSONWithFailure(c, err)
		return
	}
	ginx.JSONWithSuccess(c, resp)
	return
}

// GetPicCaptchaAnswer
// @description Get the answer code of a picture captcha with the captcha id.
// @id get-pic-captcha-answer
// @tags staff
// @summary Get the answer code of a picture captcha.
// @produce json
// @param id query string true "a captcha id."
// @security ApiKeyAuth
// @success 2000 {string} string "captcha answer code"
// @failure 1000 {error} error "Invalid Token, Invalid Captcha, Unauthorized, or other errors."
// @router /picCaptchaAnswer [GET]
func (a *Staff) GetPicCaptchaAnswer(c *gin.Context) {
	ctx := c.Request.Context()
	sessionId := ginx.SessionIdFromGinX(c)
	if sessionId.Hex() != config.C.Root.SessionId {
		ginx.JSONWithUnauthorized(c, schema.ErrUnauthorized.Error())
		return
	}
	answer := a.bll.GetPicCaptchaAnswer(ctx, c.Query("id"))
	if answer == "" {
		ginx.JSONWithInvalidCaptcha(c, schema.ErrInvalidCaptcha.Error())
		return
	}
	ginx.JSONWithSuccess(c, answer)
	return
}

// SignIn
// @description Sign in with account and password, supporting picture captcha authentication.
// @id sign-in
// @tags staff
// @summary Sign in.
// @accept json
// @produce json
// @param body body bll.SignInReq true "required attributes to sign in."
// @success 2000 {object} bll.SignInResp "sign in token and expiry time."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /signIn [PATCH]
func (a *Staff) SignIn(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.SignInReq
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

// SignOut
// @description Sign out the session account.
// @id sign-out
// @tags staff
// @summary Sign out.
// @produce json
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /signOut [PATCH]
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

// SignUp
// @description Sign up a new staff account.
// @id sign-up
// @tags staff
// @summary Sign up.
// @accept json
// @produce json
// @param body body bll.SignUpReq true "required attributes to sign up."
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /signUp [POST]
func (a *Staff) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.SignUpReq
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

// SignOff
// @description Sign off the session account, or some specific account only by root account.
// @id sign-off
// @tags staff
// @summary Sign off.
// @produce json
// @param id path string true "id of the account to sign off."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /signOff/:id [DELETE]
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

// Update
// @description Update all updatable attributes of a staff account.
// @id staff-update
// @tags staff
// @summary Update attributes of a staff.
// @accept json
// @produce json
// @param id path string true "id of the staff account to update."
// @param body body bll.StaffUpdateReq true "attributes to update."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /staff/:id [PUT]
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

// Profile
// @description Get the profile of a staff account.
// @id staff-profile
// @tags staff
// @summary Get infos of a staff account.
// @produce json
// @param id path string true "id of the staff account."
// @security ApiKeyAuth
// @success 2000 {object} bll.ProfileResp "profile of the staff."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /staff/:id [GET]
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

// List
// @description Get a list of staff accounts.
// @id staff-list
// @tags staff
// @summary Get a list of staff accounts.
// @product json
// @param query query bll.StaffListReq false "search criteria."
// @security ApiKeyAuth
// @success 2000 {object} bll.StaffListResp "staff list."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /staff [GET]
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
