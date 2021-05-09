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

func (a *Role) AddRole(ctx context.Context, creator *primitive.ObjectID, arg *AddRoleReq) error {
	newM := model.Role{
		Group:   &arg.Group,
		Name:    &arg.Name,
		Seq:     &arg.Seq,
		Icon:    &arg.Icon,
		Memo:    &arg.Memo,
		Enable:  &arg.Enable,
		Creator: creator,
	}
	if err := a.roleRepo.InsertOne(ctx, &newM); err != nil {
		return err
	}
	return nil
}
