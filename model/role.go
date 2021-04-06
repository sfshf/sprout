package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	ID      *primitive.ObjectID `bson:"_id,omitempty"`
	Name    *string             `bson:"name,omitempty"`
	Weight  *int                `bson:"weight,omitempty"`
	Memo    *string             `bson:"memo,omitempty"`
	Creator *primitive.ObjectID `bson:"creator,omitempty"`
}

// Role names.
const (
	RootRole     = "ROOT"
	OrdinaryRole = "ORDINARY"
)
