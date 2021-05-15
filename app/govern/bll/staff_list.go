package bll

import (
	"context"
	"github.com/sfshf/sprout/app/govern/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"strings"
)

type ListStaffReq struct {
	Account             string          `form:"account" binding:""`
	SignIn              bool            `form:"signIn" binding:""`
	Email               string          `form:"email" binding:""`
	Phone               string          `form:"phone" binding:""`
	Gender              string          `form:"gender" binding:""`
	Role                string          `form:"role" binding:""`
	LastSignInIp        string          `form:"lastSignInIp" binding:""`
	LastSignInTimeBegin int64           `form:"lastSignInTimeBegin" binding:"gte=0"`
	LastSignInTimeEnd   int64           `form:"lastSignInTimeEnd" binding:"gte=0"`
	OrderBy             *schema.OrderBy `form:"orderBy" binding:""`
	schema.PaginationReq
}

type StaffListElem struct {
	ID             string `json:"id,omitempty"`
	Account        string `json:"account,omitempty"`
	SignIn         bool   `json:"signIn,omitempty"`
	Email          string `json:"email,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Gender         string `json:"gender,omitempty"`
	Role           string `json:"role,omitempty"`
	SignUpAt       int64  `json:"signUpAt,omitempty"`
	LastSignInIp   string `json:"lastSignInIp,omitempty"`
	LastSignInTime int64  `json:"lastSignInTime,omitempty"`
}

type ListStaffResp struct {
	schema.PaginationResp
}

// TODO: model to schema, should use reflect?!
func (a *Staff) ListStaff(ctx context.Context, arg *ListStaffReq, sort bson.M) (*ListStaffResp, error) {
	var and bson.A
	if arg.Account != "" {
		and = append(and, bson.M{"account": arg.Account})
	}
	if arg.SignIn {
		and = append(and, bson.M{"signInToken": bson.M{"$exists": bsonx.Boolean(true)}})
		and = append(and, bson.M{"signInToken": bson.M{"$ne": ""}})
	}
	if arg.Email != "" {
		and = append(and, bson.M{"email": arg.Email})
	}
	if arg.Phone != "" {
		and = append(and, bson.M{"phone": arg.Phone})
	}
	if arg.Gender != "" {
		and = append(and, bson.M{"gender": strings.ToUpper(arg.Gender)})
	}
	if arg.Role != "" {
		and = append(and, bson.M{"role": strings.ToUpper(arg.Role)})
	}
	if arg.LastSignInIp != "" {
		and = append(and, bson.M{"lastSignInIp": arg.LastSignInIp})
	}
	if arg.LastSignInTimeBegin > 0 {
		and = append(and, bson.M{"lastSignInTime": bson.M{"$gte": primitive.DateTime(arg.LastSignInTimeBegin)}})
	}
	if arg.LastSignInTimeEnd > 0 {
		and = append(and, bson.M{"lastSignInTime": bson.M{"$lt": primitive.DateTime(arg.LastSignInTimeEnd)}})
	}
	var filter bson.M
	if len(and) > 0 {
		filter = bson.M{"$and": and}
	}
	total, err := a.staffRepo.CountByFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	opt := options.Find().SetSort(sort).SetSkip(arg.PerPage * (arg.Page - 1)).SetLimit(arg.PerPage)
	res, err := a.staffRepo.FindManyByFilter(ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	data := make([]StaffListElem, 0, len(res))
	for _, v := range res {
		elem := StaffListElem{
			ID:       v.ID.Hex(),
			Account:  *v.Account,
			SignUpAt: int64(*v.SignUpAt),
		}
		if v.Email != nil {
			elem.Email = *v.Email
		}
		if v.Phone != nil {
			elem.Phone = *v.Phone
		}
		if v.Gender != nil {
			elem.Gender = *v.Gender
		}
		if v.LastSignInIp != nil {
			elem.LastSignInIp = *v.LastSignInIp
		}
		if v.LastSignInTime != nil {
			elem.LastSignInTime = int64(*v.LastSignInTime)
		}
		data = append(data, elem)
	}
	return &ListStaffResp{
		schema.PaginationResp{
			Data:  data,
			Total: total,
		},
	}, nil
}
