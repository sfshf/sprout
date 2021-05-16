package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO need to remove associations between menu-widgets and roles, and remove casbin policies.
func (a *Menu) Evict(ctx context.Context, argId *primitive.ObjectID) error {
	return nil
}
