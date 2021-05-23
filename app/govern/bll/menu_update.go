package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateMenuReq struct {
	Name     string  `json:"name" binding:""`
	Seq      int     `json:"seq" binding:""`
	Icon     *string `json:"icon" binding:""`
	Route    string  `json:"route" binding:""`
	Memo     *string `json:"memo" binding:""`
	Show     *bool   `json:"show" binding:""`
	ParentID *string `json:"parentID" binding:""`
}

func (a *Menu) Update(ctx context.Context, objId *primitive.ObjectID, req *UpdateMenuReq) error {
	arg := &model.Menu{ID: objId}
	if req.Name != "" {
		arg.Name = &req.Name
	}
	if req.Seq > 0 {
		arg.Seq = &req.Seq
	}
	if req.Icon != nil {
		arg.Icon = req.Icon
	}
	if req.Route != "" {
		arg.Route = &req.Route
	}
	if req.Memo != nil {
		arg.Memo = req.Memo
	}
	if req.Show != nil {
		arg.Show = req.Show
	}
	if req.ParentID != nil {
		if *req.ParentID == "" {
			arg.ParentID = &primitive.ObjectID{}
		} else {
			parentId, err := primitive.ObjectIDFromHex(*req.ParentID)
			if err != nil {
				return err
			}
			arg.ParentID = &parentId
		}
	}
	return a.menuRepo.UpdateOneByID(ctx, arg)
}
