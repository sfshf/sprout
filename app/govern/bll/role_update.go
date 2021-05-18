package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateRoleReq struct {
	Group *string `json:"group" binding:"omitempty"`
	Name  string  `json:"name" binding:"omitempty"`
	Seq   *int    `json:"seq" binding:"omitempty"`
	Icon  *string `json:"icon" binding:"omitempty"`
	Memo  *string `json:"memo" binding:"omitempty"`
}

func (a *Role) Update(ctx context.Context, objId *primitive.ObjectID, req *UpdateRoleReq) error {
	arg := &model.Role{ID: objId}
	if req.Group != nil {
		arg.Group = req.Group
	}
	if req.Name != "" {
		arg.Name = &req.Name
	}
	if req.Seq != nil {
		arg.Seq = req.Seq
	}
	if req.Icon != nil {
		arg.Icon = req.Icon
	}
	if req.Memo != nil {
		arg.Memo = req.Memo
	}
	return a.roleRepo.UpdateOneByID(ctx, arg)
}
