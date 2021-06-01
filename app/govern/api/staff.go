package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sfshf/sprout/app/govern/bll"
	"github.com/sfshf/sprout/app/govern/config"
	"github.com/sfshf/sprout/app/govern/ginx"
	"github.com/sfshf/sprout/app/govern/schema"
	"github.com/sfshf/sprout/model"
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
		for _, ip := range *staff.SignInIpWhitelist {
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
	staffId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	sessionId := ginx.SessionIdFromGinX(c)
	if (staffId.Hex() != sessionId.Hex() && sessionId.Hex() != config.C.Root.SessionId) ||
		(sessionId.Hex() == config.C.Root.SessionId && staffId.Hex() == config.C.Root.SessionId) {
		ginx.JSONWithUnauthorized(c, schema.ErrUnauthorized.Error())
		return
	}
	if err := a.bll.SignOff(ctx, &staffId); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

// UpdatePassword
// @description Update the password of a staff account.
// @id staff-update-password
// @tags staff
// @summary Update the password of a staff account.
// @accept json
// @produce json
// @param id path string true "id of the staff account to update."
// @param body body bll.UpdateStaffPasswordReq true "attributes need to update the password."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /staffs/:id/password [PATCH]
func (a *Staff) UpdatePassword(c *gin.Context) {
	ctx := c.Request.Context()
	staffId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	var arg bll.UpdateStaffPasswordReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if arg.NewPassword != arg.NewPasswordRepeat {
		ginx.JSONWithInvalidArguments(c, "conflicting new password from it's repeat")
		return
	}
	if err := a.bll.UpdatePassword(ctx, &staffId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

// UpdateEmail
// @description Update the email of a staff account.
// @id staff-update-email
// @tags staff
// @summary Update the email of a staff account.
// @accept json
// @produce json
// @param id path string true "id of the staff account to update."
// @param body body bll.UpdateStaffEmailReq true "attributes need to update the email."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /staffs/:id/email [PATCH]
func (a *Staff) UpdateEmail(c *gin.Context) {
	ctx := c.Request.Context()
	staffId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	var arg bll.UpdateStaffEmailReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	// TODO email captcha.
	sessionId := ginx.SessionIdFromGinX(c)
	if arg.CaptchaFromOldEmail != "888888" && sessionId.Hex() != config.C.Root.SessionId {
		ginx.JSONWithInvalidCaptcha(c, schema.ErrInvalidCaptcha.Error())
		return
	}
	if err := a.bll.UpdateEmail(ctx, &staffId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

// UpdatePhone
// @description Update the phone of a staff account.
// @id staff-update-phone
// @tags staff
// @summary Update the phone of a staff account.
// @accept json
// @produce json
// @param id path string true "id of the staff account to update."
// @param body body bll.UpdateStaffPasswordReq true "attributes need to update the phone."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /staffs/:id/phone [PATCH]
func (a *Staff) UpdatePhone(c *gin.Context) {
	ctx := c.Request.Context()
	staffId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	var arg bll.UpdateStaffPhoneReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	// TODO phone captcha.
	sessionId := ginx.SessionIdFromGinX(c)
	if arg.CaptchaFromOldPhone != "888888" && sessionId.Hex() != config.C.Root.SessionId {
		ginx.JSONWithInvalidCaptcha(c, schema.ErrInvalidCaptcha.Error())
		return
	}
	if err := a.bll.UpdatePhone(ctx, &staffId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

// UpdateRoles
// @description Update the roles of a staff account.
// @id staff-update-roles
// @tags staff
// @summary Update the roles of a staff account.
// @accept json
// @produce json
// @param id path string true "id of the staff account to update."
// @param body body bll.UpdateStaffPasswordReq true "attributes need to update the roles."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /staffs/:id/roles [PATCH]
func (a *Staff) UpdateRoles(c *gin.Context) {
	ctx := c.Request.Context()
	staffId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	var arg bll.UpdateStaffRolesReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.UpdateRoles(ctx, &staffId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

// UpdateSignInIpWhitelist
// @description Update the sign-in ip-whitelist of a staff account.
// @id staff-update-signInIpWhitelist
// @tags staff
// @summary Update the sign-in ip-whitelist of a staff account.
// @accept json
// @produce json
// @param id path string true "id of the staff account to update."
// @param body body bll.UpdateStaffSignInIpWhitelistReq true "attributes need to update the sign-in ip-whitelist."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /staffs/:id/signInIpWhitelist [PATCH]
func (a *Staff) UpdateSignInIpWhitelist(c *gin.Context) {
	ctx := c.Request.Context()
	staffId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	var arg bll.UpdateStaffSignInIpWhitelistReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.UpdateSignInIpWhitelist(ctx, &staffId, &arg); err != nil {
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
// @success 2000 {object} bll.ProfileStaffResp "profile of the staff."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /staffs/:id [GET]
func (a *Staff) Profile(c *gin.Context) {
	ctx := c.Request.Context()
	staffId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	res, err := a.bll.Profile(ctx, &staffId)
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
// @produce json
// @param query query bll.ListStaffReq false "search criteria."
// @security ApiKeyAuth
// @success 2000 {object} bll.ListStaffResp "staff list."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /staffs [GET]
func (a *Staff) List(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.ListStaffReq
	if err := c.ShouldBindQuery(&arg); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	sort, err := schema.OrderByToBsonM(arg.OrderBy)
	if err != nil {
		ginx.JSONWithInvalidArguments(c, schema.ErrInvalidArguments.Error())
		return
	}
	res, err := a.bll.List(ctx, &arg, sort)
	if err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, res)
	return
}

// Enable
// @description Enable or disable a staff account.
// @id staff-enable
// @tags staff
// @summary Enable or disable a staff account.
// @accept json
// @produce json
// @param id path string true "id of the staff account."
// @param enable body bool true "true for enable, or false for disable."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /staffs/:id/enable [PATCH]
func (a *Staff) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.EnableStaffReq
	if err := c.ShouldBindJSON(&arg); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	staffId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.Enable(ctx, &staffId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}
