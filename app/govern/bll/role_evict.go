package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO verify validity of  multi-level roles.
func (a *Role) Evict(ctx context.Context, argId *primitive.ObjectID) error {
	arg, err := a.roleRepo.FindOneByID(ctx, argId)
	if err != nil {
		return err
	}
	_, err = a.enforcer.RemoveFilteredPolicy(0, argId.Hex())
	if err != nil {
		return err
	}
	staffIDs, err := a.enforcer.GetUsersForRole(argId.Hex())
	if err != nil {
		return err
	}
	_, err = a.enforcer.RemoveFilteredGroupingPolicy(1, argId.Hex())
	if err != nil {
		return err
	}
	for _, staffID := range staffIDs {
		staffId, err := primitive.ObjectIDFromHex(staffID)
		if err != nil {
			return err
		}
		err = a.staffRepo.EvictRole(ctx, &staffId, arg.Name)
		if err != nil {
			return err
		}
	}
	return a.roleRepo.EvictRole(ctx, argId)
}
