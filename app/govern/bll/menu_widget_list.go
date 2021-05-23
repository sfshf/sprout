package bll

import (
	"context"
	"github.com/sfshf/sprout/app/govern/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type WidgetListElem struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Seq       int    `json:"seq,omitempty"`
	Show      bool   `json:"show,omitempty"`
	Creator   string `json:"creator,omitempty"`
	Enable    bool   `json:"enable,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
}

type ListWidgetResp struct {
	schema.PaginationResp
}

func (a *Menu) ListWidget(ctx context.Context, objId *primitive.ObjectID) (*ListWidgetResp, error) {
	res, err := a.menuRepo.FindOneByFilter(ctx, bson.M{"_id": objId}, options.FindOne().SetProjection(bson.M{"widgets": bsonx.Int32(1)}))
	if err != nil {
		return nil, err
	}
	data := make([]WidgetListElem, 0, len(*res.Widgets))
	for _, v := range *res.Widgets {
		elem := WidgetListElem{
			ID:   v.ID.Hex(),
			Name: *v.Name,
			Seq:  *v.Seq,
			Show: *v.Show,
			// TODO should use creator's account.
			Creator:   v.Creator.Hex(),
			Enable:    *v.Enable,
			CreatedAt: int64(*v.CreatedAt),
		}
		if v.UpdatedAt != nil {
			elem.UpdatedAt = int64(*v.UpdatedAt)
		}
		data = append(data, elem)
	}
	return &ListWidgetResp{
		schema.PaginationResp{
			Data:  data,
			Total: int64(len(data)),
		},
	}, nil
}
