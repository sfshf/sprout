package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Role) Evict(ctx context.Context, objId *primitive.ObjectID) error {
	obj, err := a.roleRepo.FindOneAndDeleteByID(ctx, objId)
	if err != nil {
		return err
	}
	staffIDs, err := a.enforcer.GetUsersForRole(objId.Hex())
	if err != nil {
		return err
	}
	for _, staffID := range staffIDs {
		staffId, err := primitive.ObjectIDFromHex(staffID)
		if err != nil {
			return err
		}
		if _, err = a.enforcer.DeleteRoleForUser(staffId.Hex(), objId.Hex()); err != nil {
			return err
		}
		if err = a.staffRepo.EvictRole(ctx, &staffId, obj.Name); err != nil {
			return err
		}
	}
	if _, err = a.enforcer.DeleteRole(objId.Hex()); err != nil {
		return err
	}
	return nil
}
