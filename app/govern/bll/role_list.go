package bll

import (
	"context"
	"github.com/sfshf/sprout/app/govern/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ListRoleReq struct {
	Group          string          `form:"group" binding:""`
	Name           string          `form:"name" binding:""`
	Seq            int             `form:"seq" binding:"gte=0"`
	Enable         *bool           `form:"enable" binding:""`
	Creator        string          `form:"creator" binding:""`
	CreatedAtBegin int64           `form:"createdAtBegin" binding:"gte=0"`
	CreatedAtEnd   int64           `form:"createdAtEnd" binding:"gte=0"`
	OrderBy        *schema.OrderBy `form:"orderBy" binding:""`
	schema.PaginationReq
}

type RoleListElem struct {
	ID        string `json:"id,omitempty"`
	Group     string `json:"group,omitempty"`
	Name      string `json:"name,omitempty"`
	Seq       int    `json:"seq,omitempty"`
	Icon      string `json:"icon,omitempty"`
	Memo      string `json:"memo,omitempty"`
	Enable    bool   `json:"enable,omitempty"`
	Creator   string `json:"creator,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
}

type ListRoleResp struct {
	schema.PaginationResp
}

func (a *Role) ListRole(ctx context.Context, arg *ListRoleReq, sort bson.M) (*ListRoleResp, error) {
	var and bson.A
	if arg.Group != "" {
		and = append(and, bson.M{"group": arg.Group})
	}
	if arg.Name != "" {
		and = append(and, bson.M{"name": arg.Name})
	}
	if arg.Seq > 0 {
		and = append(and, bson.M{"seq": arg.Seq})
	}
	if arg.Enable != nil {
		and = append(and, bson.M{"enable": arg.Enable})
	}
	if arg.Creator != "" {
		and = append(and, bson.M{"creator": arg.Creator})
	}
	if arg.CreatedAtBegin > 0 {
		and = append(and, bson.M{"createdAt": bson.M{"$gte": primitive.DateTime(arg.CreatedAtBegin)}})
	}
	if arg.CreatedAtEnd > 0 {
		and = append(and, bson.M{"createdAt": bson.M{"$lt": primitive.DateTime(arg.CreatedAtBegin)}})
	}
	var filter bson.M
	if len(and) > 0 {
		filter = bson.M{"$and": and}
	}
	total, err := a.roleRepo.CountByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	opt := options.Find().SetSort(sort).SetSkip(arg.PerPage * (arg.Page - 1)).SetLimit(arg.PerPage)
	res, err := a.roleRepo.FindManyByFilter(ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	data := make([]RoleListElem, 0, len(res))
	for _, v := range res {
		elem := RoleListElem{
			ID:        v.ID.Hex(),
			Name:      *v.Name,
			Seq:       *v.Seq,
			Enable:    *v.Enable,
			Creator:   v.Creator.Hex(),
			CreatedAt: int64(*v.CreatedAt),
		}
		if v.Icon != nil {
			elem.Icon = *v.Icon
		}
		if v.Memo != nil {
			elem.Memo = *v.Memo
		}
		if v.UpdatedAt != nil {
			elem.UpdatedAt = int64(*v.UpdatedAt)
		}
		data = append(data, elem)
	}
	return &ListRoleResp{
		schema.PaginationResp{
			Data:  data,
			Total: total,
		},
	}, nil
}
