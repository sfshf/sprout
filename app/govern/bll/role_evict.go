package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Role) EvictRole(ctx context.Context, id *primitive.ObjectID) error {
	role, err := a.roleRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	_, err = a.enforcer.RemoveFilteredPolicy(0, id.Hex())
	if err != nil {
		return err
	}
	staffIDs, err := a.enforcer.GetUsersForRole(id.Hex())
	if err != nil {
		return err
	}
	_, err = a.enforcer.RemoveFilteredGroupingPolicy(1, id.Hex())
	if err != nil {
		return err
	}
	for _, staffID := range staffIDs {
		staffId, err := primitive.ObjectIDFromHex(staffID)
		if err != nil {
			return err
		}
		err = a.staffRepo.EvictRole(ctx, &staffId, role.Name)
		if err != nil {
			return err
		}
	}
	return a.roleRepo.EvictRole(ctx, id)
}
