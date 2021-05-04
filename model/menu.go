package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Menu struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty"`
	Name      *string             `bson:"name,omitempty"`
	Seq       *int                `bson:"seq,omitempty"`
	Icon      *string             `bson:"icon,omitempty"`
	Route     *string             `bson:"route,omitempty"`
	Memo      *string             `bson:"memo,omitempty"`
	Show      *bool               `bson:"show,omitempty"`    // true: show; false/nil: hide.
	Widgets   []Widget            `bson:"widgets,omitempty"` // buttons or input boxes.
	ParentID  *primitive.ObjectID `bson:"parentID,omitempty"`
	Creator   *primitive.ObjectID `bson:"creator,omitempty"`
	Enable    *bool               `bson:"enable,omitempty"` // true: enable; false/nil: disable.
	CreatedAt *primitive.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt *primitive.DateTime `bson:"updatedAt,omitempty"`
}

type Widget struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty"`
	Name      *string             `bson:"name,omitempty"`
	Seq       *int                `bson:"seq,omitempty"`
	Icon      *string             `bson:"icon,omitempty"`
	Api       *primitive.ObjectID `bson:"api,omitempty"`
	Memo      *string             `bson:"memo,omitempty"`
	Show      *string             `bson:"show,omitempty"` // true: show; false/nil: hide.
	ParentID  *primitive.ObjectID `bson:"parentID,omitempty"`
	Creator   *primitive.ObjectID `bson:"creator,omitempty"`
	Enable    *bool               `bson:"enable,omitempty"` // true: enable; false/nil: disable.
	CreatedAt *primitive.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt *primitive.DateTime `bson:"updatedAt,omitempty"`
}
