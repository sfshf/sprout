package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Casbin) RolesOfStaff(ctx context.Context, objId *primitive.ObjectID) ([]string, error) {
	roles, err := a.enforcer.GetRolesForUser(objId.Hex())
	if err != nil {
		return nil, err
	}
	return roles, nil
}
