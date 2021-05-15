package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

func (a *Casbin) SetRole(ctx context.Context, staffId *primitive.ObjectID, role string) error {
	if _, err := a.enforcer.AddRoleForUser(staffId.Hex(), strings.ToUpper(role)); err != nil {
		return err
	}
	roles, err := a.enforcer.GetRolesForUser(staffId.Hex())
	if err != nil {
		return err
	}
	obj := &model.Staff{
		ID:    staffId,
		Roles: &roles,
	}
	return a.staffRepo.UpdateOneByID(ctx, obj)
}
