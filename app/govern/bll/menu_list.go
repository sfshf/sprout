package bll

import (
	"context"
	"github.com/sfshf/sprout/app/govern/schema"
	"go.mongodb.org/mongo-driver/bson"
)

type ListMenuReq struct {
	Name   string `form:"name" binding:""`
	Route  string `form:"route" binding:""`
	Show   *bool  `form:"show" binding:""`
	Enable *bool  `form:"enable" binding:""`
}

type MenuListElem struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Seq       int    `json:"seq,omitempty"`
	Route     string `json:"route,omitempty"`
	Show      bool   `json:"show,omitempty"`
	ParentID  string `json:"parentID,omitempty"`
	Creator   string `json:"creator,omitempty"`
	Enable    bool   `json:"enable,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
}

type ListMenuResp struct {
	schema.PaginationResp
}

func (a *Menu) List(ctx context.Context, req *ListMenuReq, sort bson.M) (*ListMenuResp, error) {
	var and bson.A
	if req.Name != "" {
		and = append(and, bson.M{"name": req.Name})
	}
	if req.Route != "" {
		and = append(and, bson.M{"route": req.Route})
	}
	if req.Show != nil {
		and = append(and, bson.M{"show": req.Show})
	}
	if req.Enable != nil {
		and = append(and, bson.M{"enable": req.Enable})
	}
	var filter bson.M
	if len(and) > 0 {
		filter = bson.M{"$and": and}
	}
	total, err := a.menuRepo.CountByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	res, err := a.menuRepo.FindManyByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	data := make([]MenuListElem, 0, len(res))
	for _, v := range res {
		elem := MenuListElem{
			ID:       v.ID.Hex(),
			Name:     *v.Name,
			Seq:      *v.Seq,
			Route:    *v.Route,
			Show:     *v.Show,
			ParentID: v.ParentID.Hex(),
			// TODO should use the creator's account.
			Creator:   v.Creator.Hex(),
			Enable:    *v.Enable,
			CreatedAt: int64(*v.CreatedAt),
		}
		if v.UpdatedAt != nil {
			elem.UpdatedAt = int64(*v.UpdatedAt)
		}
		data = append(data, elem)
	}
	return &ListMenuResp{
		schema.PaginationResp{
			Data:  data,
			Total: total,
		},
	}, nil
}
