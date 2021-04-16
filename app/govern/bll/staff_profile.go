package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileResp struct {
	Account           string   `json:"account"`
	RealName          string   `json:"realName"`
	Email             string   `json:"email"`
	Phone             string   `json:"phone"`
	Gender            string   `json:"gender"`
	Role              string   `json:"role"`
	SignInIpWhiteList []string `json:"signInIpWhiteList"`
	SignUpAt          int64    `json:"signUpAt"`
	LastSignInIp      string   `json:"lastSignInIp"`
	LastSignInTime    int64    `json:"lastSignInTime"`
}

func (a *Staff) Profile(ctx context.Context, objId *primitive.ObjectID) (*ProfileResp, error) {
	obj, err := a.staffRepo.FindOneByID(ctx, objId)
	if err != nil {
		return nil, err
	}
	res := ProfileResp{
		Account:           *obj.Account,
		SignInIpWhiteList: obj.SignInIpWhitelist,
		SignUpAt:          int64(*obj.SignUpAt),
	}
	if obj.RealName != nil {
		res.RealName = *obj.RealName
	}
	if obj.Gender != nil {
		res.Gender = *obj.Gender
	}
	if obj.Email != nil {
		res.Email = *obj.Email
	}
	if obj.Phone != nil {
		res.Phone = *obj.Phone
	}
	if obj.LastSignInIp != nil {
		res.LastSignInIp = *obj.LastSignInIp
	}
	if obj.LastSignInTime != nil {
		res.LastSignInTime = int64(*obj.LastSignInTime)
	}
	return &res, nil
}
