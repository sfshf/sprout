package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddApiReq struct {
	Group  string `json:"group" binding:""`
	Method string `json:"method" binding:"required"`
	Path   string `json:"path" binding:"required"`
	Enable bool   `json:"enable" binding:""`
}

func (a *Api) Add(ctx context.Context, creator *primitive.ObjectID, req *AddApiReq) error {
	newM := &model.Api{
		Group:  &req.Group,
		Method: &req.Method,
		Path:   &req.Path,
		Enable: &req.Enable,
	}
	return a.apiRepo.InsertOne(ctx, newM)
}
