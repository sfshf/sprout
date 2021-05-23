package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Api) Evict(ctx context.Context, objId *primitive.ObjectID) error {
	oldObj, err := a.apiRepo.FindOneAndDeleteByID(ctx, objId)
	if err != nil {
		return err
	}
	rules := a.enforcer.GetFilteredPolicy(1, *oldObj.Path)
	needToEvicted := make([][]string, 0)
	for idx := range rules {
		if len(rules[idx]) > 2 && rules[idx][2] == *oldObj.Method {
			needToEvicted = append(needToEvicted, rules[idx])
		}
	}
	_, err = a.enforcer.RemovePolicies(needToEvicted)
	if err != nil {
		return err
	}
	// TODO evict the associations between widgets and the api, if necessary.
	return nil
}
