package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	ID          *primitive.ObjectID                          `bson:"_id,omitempty"`
	Group       *string                                      `bson:"group,omitempty"`
	Name        *string                                      `bson:"name,omitempty"`
	Seq         *int                                         `bson:"seq,omitempty"`
	Icon        *string                                      `bson:"icon,omitempty"`
	Memo        *string                                      `bson:"memo,omitempty"`
	Enable      *bool                                        `bson:"enable,omitempty"` // true: enable; false/nil: disable.
	MenuWidgets *map[primitive.ObjectID][]primitive.ObjectID `bson:"menuWidgets,omitempty"`
	Creator     *primitive.ObjectID                          `bson:"creator,omitempty"`
	CreatedAt   *primitive.DateTime                          `bson:"createdAt,omitempty"`
	UpdatedAt   *primitive.DateTime                          `bson:"updatedAt,omitempty"`
}
