package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EnableRoleReq struct {
	Enable *bool `json:"enable" binding:"required"`
}

func (a *Role) Enable(ctx context.Context, objId *primitive.ObjectID, req *EnableRoleReq) error {
	obj := &model.Role{
		ID:     objId,
		Enable: req.Enable,
	}
	return a.roleRepo.UpdateOneByID(ctx, obj)
}
