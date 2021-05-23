package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type AddWidgetReq struct {
	Name     string `json:"name" binding:"required"`
	Seq      int    `json:"seq" binding:"required"`
	Icon     string `json:"icon" binding:""`
	Api      string `json:"api" binding:"required"`
	Memo     string `json:"memo" binding:""`
	Show     bool   `json:"show" binding:""`
	ParentID string `json:"parentID" binding:""`
	Enable   bool   `json:"enable" binding:""`
}

func (a *Menu) AddWidget(ctx context.Context, creator *primitive.ObjectID, menuId *primitive.ObjectID, req *AddWidgetReq) error {
	newWidget := &model.Widget{
		ID:      model.NewObjectIDPtr(),
		Name:    &req.Name,
		Seq:     &req.Seq,
		Icon:    &req.Icon,
		Memo:    &req.Memo,
		Show:    &req.Show,
		Creator: creator,
		Enable:  &req.Enable,
	}
	if req.Api != "" {
		apiId, err := primitive.ObjectIDFromHex(req.Api)
		if err != nil {
			return err
		}
		api, err := a.apiRepo.FindOneByFilter(ctx, bson.M{"_id": apiId}, options.FindOne().SetProjection(bson.M{"_id": bsonx.Int32(1)}))
		if err != nil {
			return err
		}
		newWidget.Api = api.ID
	}
	return a.menuRepo.AddWidget(ctx, menuId, newWidget)
}
