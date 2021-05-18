package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

func (a *Casbin) SetRole(ctx context.Context, objId *primitive.ObjectID, role string) error {
	if _, err := a.enforcer.AddRoleForUser(objId.Hex(), strings.ToUpper(role)); err != nil {
		return err
	}
	roles, err := a.enforcer.GetRolesForUser(objId.Hex())
	if err != nil {
		return err
	}
	obj := &model.Staff{
		ID:    objId,
		Roles: &roles,
	}
	return a.staffRepo.UpdateOneByID(ctx, obj)
}
