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

func (a *Api) Profile(ctx context.Context, objId *primitive.ObjectID) (*ProfileApiResp, error) {
	obj, err := a.apiRepo.FindOneByID(ctx, objId)
	if err != nil {
		return nil, err
	}
	res := &ProfileApiResp{
		Group:     *obj.Group,
		Method:    *obj.Method,
		Path:      *obj.Path,
		Enable:    *obj.Enable,
		CreatedAt: int64(*obj.CreatedAt),
	}
	if one, err := a.staffRepo.FindOneByID(ctx, obj.Creator); err != nil {
		return nil, err
	} else {
		res.Creator = *one.Account
	}
	if obj.UpdatedAt != nil {
		res.UpdatedAt = int64(*obj.UpdatedAt)
	}
	return res, nil
}
