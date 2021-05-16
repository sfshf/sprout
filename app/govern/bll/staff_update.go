package bll

import (
	"context"
	"errors"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateStaffPasswordReq struct {
	OldPassword       string `json:"oldPassword" binding:"gte=8,lte=32"`
	NewPassword       string `json:"newPassword" binding:"gte=8,lte=32"`
	NewPasswordRepeat string `json:"newPasswordRepeat" binding:"gte=8,lte=32"`
}

func (a *Staff) UpdatePassword(ctx context.Context, objId *primitive.ObjectID, req *UpdateStaffPasswordReq) error {
	obj, err := a.staffRepo.FindOneByID(ctx, objId)
	if err != nil {
		return err
	}
	oldPassword := model.PasswdPtr(req.OldPassword, *obj.PasswordSalt)
	if oldPassword != obj.Password {
		return errors.New("old password wrong")
	}
	newSalt := model.NewPasswdSalt()
	arg := &model.Staff{
		ID:           obj.ID,
		Password:     model.PasswdPtr(req.NewPassword, newSalt),
		PasswordSalt: &newSalt,
	}
	return a.staffRepo.UpdateOneByID(ctx, arg)
}

type UpdateStaffEmailReq struct {
	NewEmail            string `json:"newEmail" binding:"required,email"`
	CaptchaFromOldEmail string `json:"captchaFromOldEmail" binding:"len=6"`
}

func (a *Staff) UpdateEmail(ctx context.Context, argId *primitive.ObjectID, req *UpdateStaffEmailReq) error {
	arg := &model.Staff{
		ID:    argId,
		Email: &req.NewEmail,
	}
	return a.staffRepo.UpdateOneByID(ctx, arg)
}

type UpdateStaffPhoneReq struct {
	NewPhone            string `json:"newPhone" binding:"gte=11,lte=14"`
	CaptchaFromOldPhone string `json:"captchaFromOldPhone" binding:"len=6"`
}

func (a *Staff) UpdatePhone(ctx context.Context, argId *primitive.ObjectID, req *UpdateStaffPhoneReq) error {
	arg := &model.Staff{
		ID:    argId,
		Phone: &req.NewPhone,
	}
	return a.staffRepo.UpdateOneByID(ctx, arg)
}

type UpdateStaffRolesReq struct {
	Roles *[]string `json:"roles" binding:""`
}

func (a *Staff) UpdateRoles(ctx context.Context, argId *primitive.ObjectID, req *UpdateStaffRolesReq) error {
	arg := &model.Staff{
		ID:    argId,
		Roles: req.Roles,
	}
	return a.staffRepo.UpdateOneByID(ctx, arg)
}

type UpdateStaffSignInIpWhitelistReq struct {
	SignInIpWhitelist *[]string `json:"signInIpWhitelist" binding:"dive,ip"`
}

func (a *Staff) UpdateSignInIpWhitelist(ctx context.Context, argId *primitive.ObjectID, req *UpdateStaffSignInIpWhitelistReq) error {
	arg := &model.Staff{
		ID:                argId,
		SignInIpWhitelist: req.SignInIpWhitelist,
	}
	return a.staffRepo.UpdateOneByID(ctx, arg)
}
