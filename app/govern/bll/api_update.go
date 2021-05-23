package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateApiReq struct {
	Group  *string `json:"group" binding:""`
	Method string  `json:"method" binding:""`
	Path   string  `json:"path" binding:""`
}

func (a *Api) Update(ctx context.Context, objId *primitive.ObjectID, req *UpdateApiReq) error {
	obj := &model.Api{ID: objId}
	if req.Group != nil {
		obj.Group = req.Group
	}
	var policyNeedToChange bool
	if req.Method != "" {
		obj.Method = &req.Method
		policyNeedToChange = true
	}
	if req.Path != "" {
		obj.Path = &req.Path
		policyNeedToChange = true
	}
	oldObj, err := a.apiRepo.FindOneAndUpdateByID(ctx, obj)
	if err != nil {
		return err
	}
	if policyNeedToChange {
		oldRules := a.enforcer.GetFilteredPolicy(1, *oldObj.Path)
		for idx := range oldRules {
			if len(oldRules[idx]) > 2 && oldRules[idx][2] == *oldObj.Method {
				newRule := make([]string, 0)
				newRule[0] = oldRules[idx][0]
				if obj.Path != nil {
					newRule[1] = *obj.Path
				} else {
					newRule[1] = *oldObj.Path
				}
				if obj.Method != nil {
					newRule[2] = *obj.Method
				} else {
					newRule[2] = *oldObj.Method
				}
				_, err = a.enforcer.UpdatePolicy(oldRules[idx], newRule)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
