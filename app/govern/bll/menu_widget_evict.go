package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func (a *Menu) EvictWidget(ctx context.Context, menuId *primitive.ObjectID, widgetId *primitive.ObjectID) error {
	roles, err := a.roleRepo.FindManyByFilter(
		ctx,
		bson.M{"menuWidgets": bson.M{"menuID": menuId, "widgets": widgetId}},
		options.Find().SetProjection(bson.M{"_id": bsonx.Int32(1)}))
	if err != nil {
		return err
	}
	menu, err := a.menuRepo.FindOneByFilter(ctx, bson.M{"_id": widgetId}, options.FindOne().SetProjection(bson.M{"widgets": bsonx.Int32(1)}))
	if err != nil {
		return err
	}
	for _, widget := range *menu.Widgets {
		if widget.ID.Hex() == widgetId.Hex() {
			api, err := a.apiRepo.FindOneByID(ctx, widget.Api)
			if err != nil {
				if err == mongo.ErrNoDocuments {
					break
				}
				return err
			}
			rules := make([][]string, 0)
			for _, role := range roles {
				rules = append(rules, []string{role.ID.Hex(), *api.Path, *api.Method})
			}
			_, err = a.enforcer.RemovePolicies(rules)
			if err != nil {
				return err
			}
			break
		}
	}
	return a.menuRepo.EvictWidget(ctx, menuId, widgetId)
}
