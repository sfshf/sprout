package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EnableMenuReq struct {
	Enable *bool `json:"enable" binding:"required"`
}

func (a *Menu) Enable(ctx context.Context, objId *primitive.ObjectID, req *EnableMenuReq) error {
	obj := &model.Menu{
		ID:     objId,
		Enable: req.Enable,
	}
	return a.menuRepo.UpdateOneByID(ctx, obj)
}
