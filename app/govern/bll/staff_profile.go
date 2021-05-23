package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileStaffResp struct {
	Account           string   `json:"account,omitempty"`
	RealName          string   `json:"realName,omitempty"`
	Email             string   `json:"email,omitempty"`
	Phone             string   `json:"phone,omitempty"`
	Gender            string   `json:"gender,omitempty"`
	Roles             []string `json:"roles,omitempty"`
	SignInIpWhiteList []string `json:"signInIpWhiteList,omitempty"`
	LastSignInIp      string   `json:"lastSignInIp,omitempty"`
	LastSignInTime    int64    `json:"lastSignInTime,omitempty"`
	Enable            bool     `json:"enable,omitempty"`
	SignUpAt          int64    `json:"signUpAt,omitempty"`
	UpdatedAt         int64    `json:"updatedAt,omitempty"`
}

func (a *Staff) Profile(ctx context.Context, objId *primitive.ObjectID) (*ProfileStaffResp, error) {
	obj, err := a.staffRepo.FindOneByID(ctx, objId)
	if err != nil {
		return nil, err
	}
	res := &ProfileStaffResp{
		Account:           *obj.Account,
		RealName:          *obj.RealName,
		Email:             *obj.Email,
		Phone:             *obj.Phone,
		Gender:            *obj.Gender,
		SignInIpWhiteList: *obj.SignInIpWhitelist,
		Enable:            *obj.Enable,
		SignUpAt:          int64(*obj.SignUpAt),
	}
	if obj.Roles != nil {
		res.Roles = *obj.Roles
	}
	if obj.LastSignInIp != nil {
		res.LastSignInIp = *obj.LastSignInIp
	}
	if obj.LastSignInTime != nil {
		res.LastSignInTime = int64(*obj.LastSignInTime)
	}
	if obj.UpdatedAt != nil {
		res.UpdatedAt = int64(*obj.UpdatedAt)
	}
	return res, nil
}
