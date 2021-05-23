package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type ProfileMenuResp struct {
	Name      string `json:"name,omitempty"`
	Seq       int    `json:"seq,omitempty"`
	Icon      string `json:"icon,omitempty"`
	Route     string `json:"route,omitempty"`
	Memo      string `json:"memo,omitempty"`
	Show      bool   `json:"show,omitempty"`
	ParentID  string `json:"parentID,omitempty"`
	Creator   string `json:"creator,omitempty"`
	Enable    bool   `json:"enable,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
}

func (a *Menu) Profile(ctx context.Context, objId *primitive.ObjectID) (*ProfileMenuResp, error) {
	obj, err := a.menuRepo.FindOneByFilter(ctx, bson.M{"_id": objId}, options.FindOne().SetProjection(bson.M{"widgets": bsonx.Int32(-1)}))
	if err != nil {
		return nil, err
	}
	res := &ProfileMenuResp{
		Name:      *obj.Name,
		Seq:       *obj.Seq,
		Icon:      *obj.Icon,
		Route:     *obj.Route,
		Memo:      *obj.Memo,
		Show:      *obj.Show,
		ParentID:  obj.ParentID.Hex(),
		Creator:   obj.Creator.Hex(),
		Enable:    *obj.Enable,
		CreatedAt: int64(*obj.CreatedAt),
	}
	if obj.UpdatedAt != nil {
		res.UpdatedAt = int64(*obj.UpdatedAt)
	}
	return res, nil
}
