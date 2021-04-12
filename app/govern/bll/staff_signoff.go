package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Staff) SignOff(ctx context.Context, objId *primitive.ObjectID) error {
	return a.staffRepo.DeleteOne(ctx, objId)
}
