package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
)

type SignupReq struct {
	Account           string   `json:"account" binding:"required,gte=2,lte=32"`
	Password          string   `json:"password" binding:"required,gte=8,lte=32"`
	RealName          string   `json:"realName" binding:"required,gte=2,lte=32"`
	Email             string   `json:"email" binding:"required,email"`
	Phone             string   `json:"phone" binding:"required,gte=11,lte=14"`
	Gender            string   `json:"gender" binding:"oneof=unknown male female"`
	SignInIpWhitelist []string `json:"signInIpWhitelist" binding:"dive,ip"`
	Timestamp         int64    `json:"timestamp" binding:"required,gte=0"`
}

func (a *Staff) SignUp(ctx context.Context, req *SignupReq) error {
	salt := model.NewPasswdSalt()
	newM := model.Staff{
		Account:           &req.Account,
		Password:          model.PasswdPtr(req.Password, salt),
		PasswordSalt:      &salt,
		RealName:          &req.RealName,
		Email:             &req.Email,
		Phone:             &req.Phone,
		Gender:            model.UpperStringPtr(req.Gender),
		SignInIpWhitelist: req.SignInIpWhitelist,
		SignUpAt:          model.DatetimePtr(req.Timestamp),
	}
	if err := a.staffRepo.InsertOne(ctx, &newM); err != nil {
		return err
	}
	return nil
}
