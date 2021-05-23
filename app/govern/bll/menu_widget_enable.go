package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EnableWidgetReq struct {
	Enable *bool `json:"enable" binding:"required"`
}

func (a *Menu) EnableWidget(ctx context.Context, menuId *primitive.ObjectID, widgetId *primitive.ObjectID, req *EnableWidgetReq) error {
	return a.menuRepo.UpdateOneByFilter(
		ctx,
		bson.M{"_id": menuId},
		bson.M{"$set": bson.M{"widgets.$[elem].enable": req.Enable}},
		options.Update().SetArrayFilters(
			options.ArrayFilters{Filters: bson.A{bson.M{"elem._id": widgetId}}},
		),
	)
}
