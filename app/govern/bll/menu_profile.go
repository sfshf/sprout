package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfileMenuResp struct {
}

func (a *Menu) Profile(ctx context.Context, argId *primitive.ObjectID) (*ProfileMenuResp, error) {
	return nil, nil
}
