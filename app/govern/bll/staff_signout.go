package bll

import (
	"context"
	"github.com/sfshf/sprout/app/govern/ginx"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Staff) SignOut(ctx context.Context, objId *primitive.ObjectID) error {
	_ = a.redisCache.Del(ctx, ginx.RedisKeyPrefix+objId.Hex())
	obj := &model.Staff{
		ID:          objId,
		SignInToken: model.StringPtr(""),
	}
	return a.staffRepo.UpdateOneByID(ctx, obj)
}
