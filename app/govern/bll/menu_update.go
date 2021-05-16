package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateMenuReq struct {
}

func (a *Menu) Update(ctx context.Context, argId *primitive.ObjectID, req *UpdateMenuReq) error {
	return nil
}
