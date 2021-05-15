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

func (a *Staff) ProfileStaff(ctx context.Context, argId *primitive.ObjectID) (*ProfileStaffResp, error) {
	arg, err := a.staffRepo.FindOneByID(ctx, argId)
	if err != nil {
		return nil, err
	}
	res := &ProfileStaffResp{
		Account:           *arg.Account,
		SignInIpWhiteList: *arg.SignInIpWhitelist,
		SignUpAt:          int64(*arg.SignUpAt),
	}
	if arg.RealName != nil {
		res.RealName = *arg.RealName
	}
	if arg.Email != nil {
		res.Email = *arg.Email
	}
	if arg.Phone != nil {
		res.Phone = *arg.Phone
	}
	if arg.Gender != nil {
		res.Gender = *arg.Gender
	}
	if arg.Roles != nil {
		res.Roles = *arg.Roles
	}
	if arg.LastSignInIp != nil {
		res.LastSignInIp = *arg.LastSignInIp
	}
	if arg.LastSignInTime != nil {
		res.LastSignInTime = int64(*arg.LastSignInTime)
	}
	if arg.Enable != nil {
		res.Enable = *arg.Enable
	}
	if arg.UpdatedAt != nil {
		res.UpdatedAt = int64(*arg.UpdatedAt)
	}
	return res, nil
}
