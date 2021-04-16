package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Casbin struct {
	ID    *primitive.ObjectID `bson:"_id"`
	PType *string             `bson:"pType"`
	V0    *string             `bson:"v0"`
	V1    *string             `bson:"v1"`
	V2    *string             `bson:"v2"`
	V3    *string             `bson:"v3"`
	V4    *string             `bson:"v4"`
	V5    *string             `bson:"v5"`
}

const (
	PType_P = "p" // Policy Type -- policy definition.
	PType_g = "g" // Policy Type -- role or group definition.
)

// Role names.
const (
	RootRole     = "root"
	OrdinaryRole = "ordinary"
)
