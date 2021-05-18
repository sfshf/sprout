package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EnableApiReq struct {
	Enable bool `json:"enable" binding:"required"`
}

// TODO enable/disable an api, and add/remove corresponding casbin policy.
func (a *Api) Enable(ctx context.Context, objId *primitive.ObjectID, req *EnableApiReq) error {
	obj := &model.Api{
		ID:     objId,
		Enable: &req.Enable,
	}
	return a.apiRepo.UpdateOneByID(ctx, obj)
}
