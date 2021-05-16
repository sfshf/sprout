package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileApiResp struct {
	Group     string `json:"group,omitempty"`
	Method    string `json:"method,omitempty"`
	Path      string `json:"path,omitempty"`
	Creator   string `json:"creator,omitempty"`
	Enable    bool   `json:"enable,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
}

func (a *Api) Profile(ctx context.Context, argId *primitive.ObjectID) (*ProfileApiResp, error) {
	arg, err := a.apiRepo.FindOneByID(ctx, argId)
	if err != nil {
		return nil, err
	}
	res := &ProfileApiResp{
		Group:     *arg.Group,
		Method:    *arg.Method,
		Path:      *arg.Path,
		Enable:    *arg.Enable,
		CreatedAt: int64(*arg.CreatedAt),
	}
	if one, err := a.staffRepo.FindOneByID(ctx, arg.Creator); err != nil {
		return nil, err
	} else {
		res.Creator = *one.Account
	}
	if arg.UpdatedAt != nil {
		res.UpdatedAt = int64(*arg.UpdatedAt)
	}
	return res, nil
}
