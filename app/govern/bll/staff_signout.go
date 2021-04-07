package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Staff) SignOut(ctx context.Context, id string) error {
	staffId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return a.staffRepo.SignOut(ctx, &staffId)
}
