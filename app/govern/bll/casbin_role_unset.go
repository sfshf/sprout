package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Casbin) UnsetRole(ctx context.Context, staffId *primitive.ObjectID, role string) error {
	_, err := a.enforcer.RemoveFilteredGroupingPolicy(1, role)
	if err != nil {
		return err
	}
	roles, err := a.enforcer.GetRolesForUser(staffId.Hex())
	if err != nil {
		return err
	}
	staff := &model.Staff{
		ID:    staffId,
		Roles: &roles,
	}
	return a.staffRepo.UpdateOneByID(ctx, staff)
}
