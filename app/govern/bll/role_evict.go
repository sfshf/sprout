package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO verify validity of  multi-level roles.
func (a *Role) Evict(ctx context.Context, objId *primitive.ObjectID) error {
	obj, err := a.roleRepo.FindOneByID(ctx, objId)
	if err != nil {
		return err
	}
	_, err = a.enforcer.RemoveFilteredPolicy(0, objId.Hex())
	if err != nil {
		return err
	}
	staffIDs, err := a.enforcer.GetUsersForRole(objId.Hex())
	if err != nil {
		return err
	}
	_, err = a.enforcer.RemoveFilteredGroupingPolicy(1, objId.Hex())
	if err != nil {
		return err
	}
	for _, staffID := range staffIDs {
		staffId, err := primitive.ObjectIDFromHex(staffID)
		if err != nil {
			return err
		}
		err = a.staffRepo.EvictRole(ctx, &staffId, obj.Name)
		if err != nil {
			return err
		}
	}
	return a.roleRepo.EvictRole(ctx, objId)
}
