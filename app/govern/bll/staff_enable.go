package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EnableStaffReq struct {
	Enable *bool `json:"enable" binding:"required"`
}

func (a *Staff) Enable(ctx context.Context, objId *primitive.ObjectID, req *EnableStaffReq) error {
	arg := &model.Staff{
		ID:     objId,
		Enable: req.Enable,
	}
	return a.staffRepo.UpdateOneByID(ctx, arg)
}
