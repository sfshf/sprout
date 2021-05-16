package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddMenuReq struct {
	Name     string    `json:"name" binding:"required"`
	Seq      int       `json:"seq" binding:"required"`
	Icon     string    `json:"icon" binding:""`
	Route    string    `json:"route" binding:"required"`
	Memo     string    `json:"memo" binding:""`
	Show     bool      `json:"show" binding:""`
	ParentID string    `json:"parentID" binding:""`
	Enable   bool      `json:"enable" binding:""`
	Widgets  []*Widget `json:"widgets" binding:""`
}

type Widget struct {
	Name     string `json:"name" binding:"required"`
	Seq      int    `json:"seq" binding:"required"`
	Icon     string `json:"icon" binding:""`
	Api      string `json:"api" binding:"required"`
	Memo     string `json:"memo" binding:""`
	Show     bool   `json:"show" binding:""`
	ParentID string `json:"parentID" binding:""`
	Enable   bool   `json:"enable" binding:""`
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
			if one.ParentID != "" {
				parentId, err := primitive.ObjectIDFromHex(one.ParentID)
				if err != nil {
					return err
				}
				widget.ParentID = &parentId
			}
			if one.Api != "" {
				apiId, err := primitive.ObjectIDFromHex(one.Api)
				if err != nil {
					return err
				}
				widget.Api = &apiId
			}
			widgets = append(widgets, widget)
		}
		newM.Widgets = &widgets
	}
	return a.menuRepo.InsertOne(ctx, newM)
}
