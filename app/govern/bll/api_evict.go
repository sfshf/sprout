package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Api) Evict(ctx context.Context, argId *primitive.ObjectID) error {
	return a.apiRepo.DeleteOne(ctx, argId)
}
