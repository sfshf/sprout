package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Casbin) RolesOfStaff(ctx context.Context, id *primitive.ObjectID) ([]string, error) {
	roles, err := a.enforcer.GetRolesForUser(id.Hex())
	if err != nil {
		return nil, err
	}
	return roles, nil
}
