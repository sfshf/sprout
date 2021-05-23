package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type AuthorizeReq struct {
	MenuWidgetMap map[string][]string `json:"menuWidgetMap" binding:"required"`
}

func (a *Role) Authorize(ctx context.Context, objId *primitive.ObjectID, req *AuthorizeReq) error {
	role, err := a.roleRepo.FindOneByID(ctx, objId)
	if err != nil {
		return err
	}
	menuWidgets := make([]model.MenuWidgets, len(req.MenuWidgetMap))
	var apiIds []*primitive.ObjectID
	for menuID, widgetIDs := range req.MenuWidgetMap {
		menuId, err := primitive.ObjectIDFromHex(menuID)
		if err != nil {
			return err
		}
		menu, err := a.menuRepo.FindOneByFilter(ctx, bson.M{"_id": menuId}, options.FindOne().SetProjection(bson.M{"widgets._id": bsonx.Int32(1), "widgets.api": bsonx.Int32(1)}))
		if err != nil {
			return err
		}
		widgets := make([]*primitive.ObjectID, 0, len(*menu.Widgets))
		for _, widget := range *menu.Widgets {
			for _, widgetID := range widgetIDs {
				if widget.ID.Hex() == widgetID {
					apiIds = append(apiIds, widget.Api)
				}
			}
			widgets = append(widgets, widget.ID)
		}
		menuWidgets = append(menuWidgets, model.MenuWidgets{
			MenuID:  &menuId,
			Widgets: widgets,
		})
	}
	if _, err = a.enforcer.DeleteRole(role.ID.Hex()); err != nil {
		return err
	}
	for _, apiId := range apiIds {
		api, err := a.apiRepo.FindOneByID(ctx, apiId)
		if err != nil {
			return err
		}
		if *api.Enable {
			_, err = a.enforcer.AddPolicy(
				objId.Hex(),
				*api.Path,
				*api.Method,
			)
			if err != nil {
				return err
			}
		}
	}
	arg := &model.Role{
		ID:          objId,
		MenuWidgets: &menuWidgets,
	}
	return a.roleRepo.UpdateOneByID(ctx, arg)
}
