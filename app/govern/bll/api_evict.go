package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO evict an api, and remove corresponding casbin policy.
func (a *Api) Evict(ctx context.Context, objId *primitive.ObjectID) error {
	return a.apiRepo.DeleteOne(ctx, objId)
}
