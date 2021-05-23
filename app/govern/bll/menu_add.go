package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type AddMenuReq struct {
	Name     string       `json:"name" binding:"required"`
	Seq      int          `json:"seq" binding:"required"`
	Icon     string       `json:"icon" binding:""`
	Route    string       `json:"route" binding:"required"`
	Memo     string       `json:"memo" binding:""`
	Show     bool         `json:"show" binding:""`
	ParentID string       `json:"parentID" binding:""`
	Enable   bool         `json:"enable" binding:""`
	Widgets  []*AddWidget `json:"widgets" binding:""`
}

type AddWidget struct {
	Name   string `json:"name" binding:"required"`
	Seq    int    `json:"seq" binding:"required"`
	Icon   string `json:"icon" binding:""`
	Api    string `json:"api" binding:"required"`
	Memo   string `json:"memo" binding:""`
	Show   bool   `json:"show" binding:""`
	Enable bool   `json:"enable" binding:""`
}

func (a *Menu) Add(ctx context.Context, creator *primitive.ObjectID, req *AddMenuReq) error {
	newM := &model.Menu{
		Name:    &req.Name,
		Seq:     &req.Seq,
		Icon:    &req.Icon,
		Route:   &req.Route,
		Memo:    &req.Memo,
		Show:    &req.Show,
		Creator: creator,
		Enable:  &req.Enable,
	}
	if req.ParentID != "" {
		parentId, err := primitive.ObjectIDFromHex(req.ParentID)
		if err != nil {
			return err
		}
		newM.ParentID = &parentId
	}
	var widgets []*model.Widget
	if req.Widgets != nil {
		for _, one := range req.Widgets {
			widget := &model.Widget{
				ID:      model.NewObjectIDPtr(),
				Name:    &one.Name,
				Seq:     &one.Seq,
				Icon:    &one.Icon,
				Memo:    &one.Memo,
				Show:    &one.Show,
				Creator: creator,
				Enable:  &one.Enable,
			}
			if one.Api != "" {
				apiId, err := primitive.ObjectIDFromHex(one.Api)
				if err != nil {
					return err
				}
				api, err := a.apiRepo.FindOneByFilter(ctx, bson.M{"_id": apiId}, options.FindOne().SetProjection(bson.M{"_id": bsonx.Int32(1)}))
				if err != nil {
					return err
				}
				widget.Api = api.ID
			}
			widgets = append(widgets, widget)
		}
		newM.Widgets = &widgets
	}
	return a.menuRepo.InsertOne(ctx, newM)
}
