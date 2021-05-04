package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Api struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty"`
	Group     *string             `bson:"group,omitempty"`
	Method    *string             `bson:"method,omitempty"`
	Path      *string             `bson:"path,omitempty"`
	Creator   *primitive.ObjectID `bson:"creator,omitempty"`
	Enable    *bool               `bson:"enable,omitempty"`
	CreatedAt *primitive.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt *primitive.DateTime `bson:"updatedAt,omitempty"`
}
