package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type EnableApiReq struct {
	Enable *bool `json:"enable" binding:"required"`
}

func (a *Api) Enable(ctx context.Context, objId *primitive.ObjectID, req *EnableApiReq) error {
	arg := &model.Api{
		ID:     objId,
		Enable: req.Enable,
	}
	api, err := a.apiRepo.FindOneAndUpdateByID(ctx, arg)
	if err != nil {
		return err
	}
	if !*req.Enable {
		rules := a.enforcer.GetFilteredPolicy(1, *api.Path)
		needToEvicted := make([][]string, 0)
		for idx := range rules {
			if len(rules[idx]) > 2 && rules[idx][2] == *api.Method {
				needToEvicted = append(needToEvicted, rules[idx])
			}
		}
		_, err = a.enforcer.RemovePolicies(needToEvicted)
		if err != nil {
			return err
		}
	} else {
		menus, err := a.menuRepo.FindManyByFilter(
			ctx,
			bson.M{"widgets": bson.M{"api": api.ID}},
			options.Find().SetProjection(bson.M{
				"_id":         bsonx.Int32(1),
				"widgets._id": bsonx.Int32(1),
				"widgets.api": bsonx.Int32(1),
			}),
		)
		if err != nil {
			return err
		}
		for idx := range menus {
			for _, widget := range *menus[idx].Widgets {
				if widget.Api.Hex() == api.ID.Hex() {
					roles, err := a.roleRepo.FindManyByFilter(ctx, bson.M{"menuWidgets": bson.M{"menuID": menus[idx].ID, "widgets": widget.ID}}, options.Find().SetProjection(bson.M{"_id": bsonx.Int32(1)}))
					if err != nil {
						return err
					}
					for i := range roles {
						_, err = a.enforcer.AddPolicy(roles[i].ID.Hex(), *api.Path, *api.Method)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}
