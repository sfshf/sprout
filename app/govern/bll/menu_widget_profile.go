package bll

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type ProfileWidgetResp struct {
	Name string `json:"name,omitempty"`
	Seq  int    `json:"seq,omitempty"`
	Icon string `json:"icon,omitempty"`
	Api  struct {
		ID     string `json:"id,omitempty"`
		Path   string `json:"path,omitempty"`
		Method string `json:"method,omitempty"`
	} `json:"api,omitempty"`
	Memo      string `json:"memo,omitempty"`
	Show      bool   `json:"show,omitempty"`
	Creator   string `json:"creator,omitempty"`
	Enable    bool   `json:"enable,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
}

func (a *Menu) ProfileWidget(ctx context.Context, menuId *primitive.ObjectID, widgetId *primitive.ObjectID) (*ProfileWidgetResp, error) {
	menu, err := a.menuRepo.FindOneByFilter(ctx, bson.M{"_id": menuId}, options.FindOne().SetProjection(bson.M{"widgets": bsonx.Int32(1)}))
	if err != nil {
		return nil, err
	}
	for _, widget := range *menu.Widgets {
		if widget.ID.Hex() == widgetId.Hex() {
			res := &ProfileWidgetResp{
				Name:      *widget.Name,
				Seq:       *widget.Seq,
				Icon:      *widget.Icon,
				Memo:      *widget.Memo,
				Show:      *widget.Show,
				Creator:   widget.Creator.Hex(),
				Enable:    *widget.Enable,
				CreatedAt: int64(*widget.CreatedAt),
			}
			if widget.UpdatedAt != nil {
				res.UpdatedAt = int64(*widget.UpdatedAt)
			}
			api, err := a.apiRepo.FindOneByID(ctx, widget.Api)
			if err != nil {
				return nil, err
			}
			res.Api.ID = api.ID.Hex()
			res.Api.Path = *api.Path
			res.Api.Method = *api.Method
			return res, nil
		}
	}
	return nil, errors.New("invalid widget id")
}
