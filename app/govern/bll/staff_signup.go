package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
)

type SignUpReq struct {
	Account           string   `json:"account" binding:"gte=2,lte=32"`
	Password          string   `json:"password" binding:"gte=8,lte=32"`
	RealName          string   `json:"realName" binding:"gte=2,lte=32"`
	Email             string   `json:"email" binding:"required,email"`
	Phone             string   `json:"phone" binding:"gte=11,lte=14"`
	Gender            string   `json:"gender" binding:"oneof=unknown male female"`
	SignInIpWhitelist []string `json:"signInIpWhitelist" binding:"dive,ip"`
	Timestamp         int64    `json:"timestamp" binding:"required,gte=0"`
}

func (a *Staff) SignUp(ctx context.Context, req *SignUpReq) error {
	salt := model.NewPasswdSalt()
	newM := &model.Staff{
		Account:           &req.Account,
		Password:          model.PasswdPtr(req.Password, salt),
		PasswordSalt:      &salt,
		RealName:          &req.RealName,
		Email:             &req.Email,
		Phone:             &req.Phone,
		Gender:            model.UpperStringPtr(req.Gender),
		SignInIpWhitelist: &req.SignInIpWhitelist,
		Enable:            model.BoolPtr(true),
		SignUpAt:          model.DatetimePtr(req.Timestamp),
	}
	return a.staffRepo.InsertOne(ctx, newM)
}
