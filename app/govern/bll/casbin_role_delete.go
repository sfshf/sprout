package bll

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Casbin) DeleteRole(ctx context.Context, role string) error {
	_, err := a.enforcer.RemoveFilteredPolicy(0, role)
	if err != nil {
		return err
	}
	ids, err := a.enforcer.GetUsersForRole(role)
	if err != nil {
		return err
	}
	_, err = a.enforcer.RemoveFilteredGroupingPolicy(1, role)
	if err != nil {
		return err
	}
	for _, id := range ids {
		staffId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}
		roles, err := a.enforcer.GetRolesForUser(staffId.Hex())
		if err != nil {
			return err
		}
		staff := &model.Staff{
			ID:    &staffId,
			Roles: &roles,
		}
		if err = a.staffRepo.UpdateOneByID(ctx, staff); err != nil {
			return err
		}
	}
	return nil
}
