package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Casbin) StaffsOfRole(ctx context.Context, role string) ([][2]string, error) {
	ids, err := a.enforcer.GetUsersForRole(role)
	if err != nil {
		return nil, err
	}
	names := make([][2]string, 0, len(ids))
	for _, id := range ids {
		staffId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		one, err := a.staffRepo.FindOneByID(ctx, &staffId)
		if err != nil {
			return nil, err
		}
		names = append(names, [2]string{*one.Account, id})
	}
	return names, nil
}
