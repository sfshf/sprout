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

// TODO Update the method or path of an api, and update corresponding casbin policies.
func (a *Api) Update(ctx context.Context, objId *primitive.ObjectID, req *UpdateApiReq) error {
	obj := &model.Api{ID: objId}
	if req.Group != nil {
		obj.Group = req.Group
	}
	if req.Method != "" {
		obj.Method = &req.Method
	}
	if req.Path != "" {
		obj.Path = &req.Path
	}
	return a.apiRepo.UpdateOneByID(ctx, obj)
}
