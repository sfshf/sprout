package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StaffUpdateReq struct {
	Password          string   `json:"password" binding:"gte=8,lte=32"`
	Email             string   `json:"email" binding:"email"`
	Phone             string   `json:"phone" binding:"gte=11,lte=14"`
	SignInIpWhitelist []string `json:"signInIpWhitelist" binding:"dive,ip"`
}

func (a *Staff) Update(ctx context.Context, objId *primitive.ObjectID, req *StaffUpdateReq) error {
	obj := &model.Staff{ID: objId}
	if req.Password != "" {
		salt := model.NewPasswdSalt()
		obj.Password = model.PasswdPtr(req.Password, salt)
		obj.PasswordSalt = &salt
	}
	if req.Email != "" {
		obj.Email = &req.Email
	}
	if req.Phone != "" {
		obj.Phone = &req.Phone
	}
	if req.SignInIpWhitelist != nil {
		obj.SignInIpWhitelist = req.SignInIpWhitelist
	}
	return a.staffRepo.UpdateOne(ctx, obj)
}
