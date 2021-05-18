package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthorizeReq struct {
	MenuWidgetMap map[string][]string `json:"menuWidgetMap" binding:"required"`
}

func (a *Role) Authorize(ctx context.Context, objId *primitive.ObjectID, req *AuthorizeReq) error {
	_, err := a.roleRepo.FindOneByID(ctx, objId)
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
		// TODO suggest to use projection-query function.
		menu, err := a.menuRepo.FindOneByID(ctx, &menuId)
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
	// TODO evict obsolete policies when necessary.
	for _, apiId := range apiIds {
		api, err := a.apiRepo.FindOneByID(ctx, apiId)
		if err != nil {
			return err
		}
		if *api.Enable {
			_, err = a.enforcer.AddPolicy(
				objId.Hex(),
				api.Path,
				api.Method,
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
