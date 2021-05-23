package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func (a *Menu) Evict(ctx context.Context, objId *primitive.ObjectID) error {
	obj, err := a.menuRepo.FindOneAndDeleteByID(ctx, objId)
	if err != nil {
		return err
	}
	roles, err := a.roleRepo.FindManyByFilter(ctx, bson.M{"menuWidgets": bson.M{"menuID": obj.ID}}, options.Find().SetProjection(bson.M{"_id": bsonx.Int32(1), "menuWidgets": bsonx.Int32(1)}))
	if err != nil {
		return err
	}
	for _, role := range roles {
		var menuWidgets []model.MenuWidgets
		for _, menuWidget := range *role.MenuWidgets {
			if menuWidget.MenuID.Hex() != obj.ID.Hex() {
				menuWidgets = append(menuWidgets, menuWidget)
			}
		}
		update := &model.Role{
			ID:          role.ID,
			MenuWidgets: &menuWidgets,
		}
		if err := a.roleRepo.UpdateOneByID(ctx, update); err != nil {
			return err
		}
	}
	var apis []*model.Api
	for _, widget := range *obj.Widgets {
		api, err := a.apiRepo.FindOneByID(ctx, widget.Api)
		if err != nil {
			return err
		}
		apis = append(apis, api)
	}
	var rules [][]string
	for _, api := range apis {
		for _, role := range roles {
			rules = append(rules, []string{role.ID.Hex(), *api.Path, *api.Method})
		}
	}
	_, err = a.enforcer.RemovePolicies(rules)
	if err != nil {
		return err
	}
	return nil
}
