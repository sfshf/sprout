package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO need to remove associations between menu-widgets and roles, and remove casbin policies.
func (a *Menu) Evict(ctx context.Context, objId *primitive.ObjectID) error {
	obj, err := a.menuRepo.FindOneByID(ctx, objId)
	if err != nil {
		return err
	}
	roles, err := a.roleRepo.FindManyByFilter(ctx, bson.M{"menuWidgets.menuID": obj.ID})
	if err != nil {
		return err
	}
	for _ = range roles {
		// TODO evict role-menu associations.
	}
	//if err = a.roleRepo.EvictMenu(ctx, obj.ID); err != nil {
	//	return err
	//}
	var apis []*model.Api
	for _, widget := range *obj.Widgets {
		api, err := a.apiRepo.FindOneByID(ctx, widget.Api)
		if err != nil {
			return err
		}
		apis = append(apis, api)
	}

	return nil
}
