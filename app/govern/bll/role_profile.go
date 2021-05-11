package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileRoleResp struct {
	Group     string `json:"group,omitempty"`
	Name      string `json:"name,omitempty"`
	Seq       int    `json:"seq,omitempty"`
	Icon      string `json:"icon,omitempty"`
	Memo      string `json:"memo,omitempty"`
	Enable    bool   `json:"enable,omitempty"`
	Creator   string `json:"creator,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
}

func (a *Role) ProfileRole(ctx context.Context, argId *primitive.ObjectID) (*ProfileRoleResp, error) {
	arg, err := a.roleRepo.FindOneByID(ctx, argId)
	if err != nil {
		return nil, err
	}
	res := &ProfileRoleResp{
		Name:      *arg.Name,
		CreatedAt: int64(*arg.CreatedAt),
	}
	if arg.Group != nil {
		res.Group = *arg.Group
	}
	if arg.Seq != nil {
		res.Seq = *arg.Seq
	}
	if arg.Icon != nil {
		res.Icon = *arg.Icon
	}
	if arg.Memo != nil {
		res.Memo = *arg.Memo
	}
	if arg.Enable != nil {
		res.Enable = *arg.Enable
	}
	if arg.Creator != nil {
		if one, err := a.staffRepo.FindOneByID(ctx, arg.Creator); err != nil {
			return nil, err
		} else {
			res.Creator = *one.Account
		}
	}
	if arg.UpdatedAt != nil {
		res.UpdatedAt = int64(*arg.UpdatedAt)
	}
	return res, nil
}
