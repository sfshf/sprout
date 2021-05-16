package bll

import (
	"context"
	"github.com/sfshf/sprout/app/govern/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ListApiReq struct {
	Group          string          `form:"group" binding:""`
	Method         string          `form:"method" binding:""`
	Path           string          `form:"path" binding:""`
	Enable         *bool           `form:"enable" binding:""`
	Creator        string          `form:"creator" binding:""`
	CreatedAtBegin int64           `form:"createdAtBegin" binding:""`
	CreatedAtEnd   int64           `form:"createdAtEnd" binding:""`
	OrderBy        *schema.OrderBy `form:"orderBy" binding:""`
	schema.PaginationReq
}

type ApiListElem struct {
	ID        string `json:"id,omitempty"`
	Group     string `json:"group,omitempty"`
	Method    string `json:"method,omitempty"`
	Path      string `json:"path,omitempty"`
	Creator   string `json:"creator,omitempty"`
	Enable    bool   `json:"enable,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
}

type ListApiResp struct {
	schema.PaginationResp
}

func (a *Api) List(ctx context.Context, req *ListApiReq, sort bson.M) (*ListApiResp, error) {
	var and bson.A
	if req.Group != "" {
		and = append(and, bson.M{"group": req.Group})
	}
	if req.Method != "" {
		and = append(and, bson.M{"method": req.Method})
	}
	if req.Path != "" {
		and = append(and, bson.M{"path": req.Path})
	}
	if req.Enable != nil {
		and = append(and, bson.M{"enable": req.Enable})
	}
	// TODO should use the creator's account.
	if req.Creator != "" {
		and = append(and, bson.M{"creator": req.Creator})
	}
	if req.CreatedAtBegin > 0 {
		and = append(and, bson.M{"createdAt": bson.M{"$gte": primitive.DateTime(req.CreatedAtBegin)}})
	}
	if req.CreatedAtEnd > 0 {
		and = append(and, bson.M{"createdAt": bson.M{"$lt": primitive.DateTime(req.CreatedAtBegin)}})
	}
	var filter bson.M
	if len(and) > 0 {
		filter = bson.M{"$and": and}
	}
	total, err := a.apiRepo.CountByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	opt := options.Find().SetSort(sort).SetSkip(req.PerPage * (req.Page - 1)).SetLimit(req.PerPage)
	res, err := a.apiRepo.FindManyByFilter(ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	data := make([]ApiListElem, 0, len(res))
	for _, v := range res {
		elem := ApiListElem{
			ID:     v.ID.Hex(),
			Group:  *v.Group,
			Method: *v.Method,
			Path:   *v.Path,
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
	return &ListApiResp{
		schema.PaginationResp{
			Data:  data,
			Total: total,
		},
	}, nil
}
