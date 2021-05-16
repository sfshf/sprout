package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddRoleReq struct {
	Group  string `json:"group" binding:""`
	Name   string `json:"name" binding:"required"`
	Seq    int    `json:"seq" binding:"required"`
	Icon   string `json:"icon" binding:""`
	Memo   string `json:"memo" binding:""`
	Enable bool   `json:"enable" binding:""`
}

func (a *Role) Add(ctx context.Context, creator *primitive.ObjectID, req *AddRoleReq) error {
	newM := &model.Role{
		Group:   &req.Group,
		Name:    &req.Name,
		Seq:     &req.Seq,
		Icon:    &req.Icon,
		Memo:    &req.Memo,
		Enable:  &req.Enable,
		Creator: creator,
	}
	return a.roleRepo.InsertOne(ctx, newM)
}
