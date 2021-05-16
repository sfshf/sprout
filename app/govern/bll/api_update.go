package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateApiReq struct {
	Group  *string `json:"group" binding:""`
	Method string  `json:"method" binding:""`
	Path   string  `json:"path" binding:""`
}

func (a *Api) Update(ctx context.Context, argId *primitive.ObjectID, req *UpdateApiReq) error {
	arg := &model.Api{ID: argId}
	if req.Group != nil {
		arg.Group = req.Group
	}
	if req.Method != "" {
		arg.Method = &req.Method
	}
	if req.Path != "" {
		arg.Path = &req.Path
	}
	return a.apiRepo.UpdateOneByID(ctx, arg)
}
