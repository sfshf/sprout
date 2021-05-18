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
		SignInIpWhiteList: *obj.SignInIpWhitelist,
		SignUpAt:          int64(*obj.SignUpAt),
	}
	if obj.RealName != nil {
		res.RealName = *obj.RealName
	}
	if obj.Email != nil {
		res.Email = *obj.Email
	}
	if obj.Phone != nil {
		res.Phone = *obj.Phone
	}
	if obj.Gender != nil {
		res.Gender = *obj.Gender
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
	if obj.Enable != nil {
		res.Enable = *obj.Enable
	}
	if obj.UpdatedAt != nil {
		res.UpdatedAt = int64(*obj.UpdatedAt)
	}
	return res, nil
}
