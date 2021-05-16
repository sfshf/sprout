package bll

import (
	"context"
	"github.com/sfshf/sprout/app/govern/schema"
	"go.mongodb.org/mongo-driver/bson"
)

type ListMenuReq struct {
	OrderBy *schema.OrderBy `form:"orderBy" binding:""`
	schema.PaginationReq
}

type MenuListElem struct {
	ID string `json:"id,omitempty"`
}

type ListMenuResp struct {
	schema.PaginationResp
}

func (a *Menu) List(ctx context.Context, req *ListMenuReq, sort bson.M) (*ListMenuResp, error) {
	return nil, nil
}
