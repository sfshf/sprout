package bll

import (
	"context"
	"github.com/sfshf/sprout/app/govern/schema"
	"go.mongodb.org/mongo-driver/bson"
)

type ListRoleReq struct {
	OrderBy *schema.OrderBy `form:"orderBy" binding:""`
	schema.PaginationReq
}

type RoleListElem struct {
}

type ListRoleResp struct {
	schema.PaginationResp
}

func (a *Role) ListRole(ctx context.Context, arg *ListRoleReq, sort bson.M) (*ListRoleResp, error) {
	return nil, nil
}
