package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateRoleReq struct {
	Group string `json:"group" binding:"omitempty"`
	Name  string `json:"name" binding:"omitempty"`
	Seq   int    `json:"seq" binding:"omitempty"`
	Icon  string `json:"icon" binding:"omitempty"`
	Memo  string `json:"memo" binding:"omitempty"`
}

func (a *Role) UpdateRole(ctx context.Context, argId *primitive.ObjectID, req *UpdateRoleReq) error {
	arg := &model.Role{ID: argId}
	if req.Group != "" {
		arg.Group = &req.Group
	}
	if req.Name != "" {
		arg.Name = &req.Name
	}
	if req.Seq != 0 {
		arg.Seq = &req.Seq
	}
	if req.Icon != "" {
		arg.Icon = &req.Icon
	}
	if req.Memo != "" {
		arg.Memo = &req.Memo
	}
	return a.roleRepo.UpdateOneByID(ctx, arg)
}
