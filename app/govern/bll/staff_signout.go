package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Staff) SignOut(ctx context.Context, objId *primitive.ObjectID) error {
	obj := &model.Staff{
		ID:          objId,
		SignInToken: model.StringPtr(""),
	}
	return a.staffRepo.UpdateOne(ctx, obj)
}
