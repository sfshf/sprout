package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateStaffReq struct {
	Password          string   `json:"password" binding:"omitempty,gte=8,lte=32"`
	Email             string   `json:"email" binding:"omitempty,email"`
	Phone             string   `json:"phone" binding:"omitempty,gte=11,lte=14"`
	SignInIpWhitelist []string `json:"signInIpWhitelist" binding:"omitempty,dive,ip"`
}

func (a *Staff) UpdateStaff(ctx context.Context, argId *primitive.ObjectID, req *UpdateStaffReq) error {
	arg := &model.Staff{ID: argId}
	if req.Password != "" {
		salt := model.NewPasswdSalt()
		arg.Password = model.PasswdPtr(req.Password, salt)
		arg.PasswordSalt = &salt
	}
	if req.Email != "" {
		arg.Email = &req.Email
	}
	if req.Phone != "" {
		arg.Phone = &req.Phone
	}
	if req.SignInIpWhitelist != nil {
		arg.SignInIpWhitelist = req.SignInIpWhitelist
	}
	return a.staffRepo.UpdateOneByID(ctx, arg)
}
