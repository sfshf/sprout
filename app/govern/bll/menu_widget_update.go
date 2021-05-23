package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type UpdateWidgetReq struct {
	Name string  `json:"name" binding:""`
	Seq  int     `json:"seq" binding:""`
	Icon *string `json:"icon" binding:""`
	Api  string  `json:"api" binding:""`
	Memo *string `json:"memo" binding:""`
	Show *bool   `json:"show" binding:""`
}

func (a *Menu) UpdateWidget(ctx context.Context, menuId *primitive.ObjectID, widgetId *primitive.ObjectID, req *UpdateWidgetReq) error {
	var sets bson.M
	if req.Name != "" {
		sets["widgets.$[elem].name"] = req.Name
	}
	if req.Seq > 0 {
		sets["widgets.$[elem].seq"] = req.Seq
	}
	if req.Icon != nil {
		sets["widgets.$[elem].icon"] = req.Icon
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
		sets["widgets.$[elem].api"] = api.ID
	}
	if req.Memo != nil {
		sets["widgets.$[elem].memo"] = req.Memo
	}
	if req.Show != nil {
		sets["widgets.$[elem].show"] = req.Show
	}
	return a.menuRepo.UpdateOneByFilter(
		ctx,
		bson.M{"_id": menuId},
		bson.M{"$set": sets},
		options.Update().SetArrayFilters(
			options.ArrayFilters{Filters: bson.A{bson.M{"elem._id": widgetId}}},
		),
	)
}
