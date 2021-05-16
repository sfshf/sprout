package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EnableMenuReq struct {
	Enable bool `json:"enable" binding:"required"`
}

func (a *Menu) Enable(ctx context.Context, objId *primitive.ObjectID, req *EnableMenuReq) error {
	return nil
}
