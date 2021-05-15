package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AllocateAuthorityReq struct {
	MenuWidgetMap map[string][]string `json:"menuWidgetMap" binding:"required"`
}

func (a *Role) AllocateAuthority(ctx context.Context, roleId *primitive.ObjectID, arg *AllocateAuthorityReq) error {
	_, err := a.roleRepo.FindOneByID(ctx, roleId)
	if err != nil {
		return err
	}
	var apiIds []*primitive.ObjectID
	for menuID, widgetIDs := range arg.MenuWidgetMap {
		menuId, err := primitive.ObjectIDFromHex(menuID)
		if err != nil {
			return err
		}
		menu, err := a.menuRepo.FindByID(ctx, &menuId)
		if err != nil {
			return err
		}
		for _, widget := range *menu.Widgets {
			for _, widgetID := range widgetIDs {
				if widget.ID.Hex() == widgetID {
					apiIds = append(apiIds, widget.Api)
				}
			}
		}
	}
	for _, apiId := range apiIds {
		api, err := a.apiRepo.FindByID(ctx, apiId)
		if err != nil {
			return err
		}
		if *api.Enable {
			_, err = a.enforcer.AddPolicy(
				roleId.Hex(),
				api.Path,
				api.Method,
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
