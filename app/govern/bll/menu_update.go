package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateMenuReq struct {
}

func (a *Menu) Update(ctx context.Context, objId *primitive.ObjectID, req *UpdateMenuReq) error {
	return nil
}
