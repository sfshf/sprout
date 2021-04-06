package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Casbin struct {
	ID    primitive.ObjectID `bson:"_id"`
	Ptype *string            `bson:"ptype"`
	V0    *string            `bson:"v0"`
	V1    *string            `bson:"v1"`
	V2    *string            `bson:"v2"`
	V3    *string            `bson:"v3"`
	V4    *string            `bson:"v4"`
	V5    *string            `bson:"v5"`
}
