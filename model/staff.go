package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	PasswdSalt = "sprout:v1:passwd:"
)

type Staff struct {
	ID                *primitive.ObjectID `bson:"_id,omitempty"`
	Account           *string             `bson:"account,omitempty"`
	Password          *string             `bson:"password,omitempty"`
	PasswordSalt      *string             `bson:"passwordSalt,omitempty"`
	RealName          *string             `bson:"realName,omitempty"`
	Email             *string             `bson:"email,omitempty"`
	Phone             *string             `bson:"phone,omitempty"`
	Gender            *string             `bson:"gender,omitempty"`
	Roles             []string            `bson:"roles,omitempty"` // a role-name slice.
	SignInIpWhitelist []string            `bson:"signInIpWhitelist,omitempty"`
	SignUpAt          *primitive.DateTime `bson:"signUpAt,omitempty"`
	SignInToken       *string             `bson:"signInToken,omitempty"`
	LastSignInIp      *string             `bson:"lastSignInIp,omitempty"`
	LastSignInTime    *primitive.DateTime `bson:"lastSignInTime,omitempty"`
	Enable            *bool               `bson:"enable,omitempty"` // true: enable; false/nil: disable.
	CreatedAt         *primitive.DateTime `bson:"createdAt,omitempty"`
	UpdatedAt         *primitive.DateTime `bson:"updatedAt,omitempty"`
}

// Genders.
const (
	UnknownGender = "UNKNOWN"
	MaleGender    = "MALE"
	FemaleGender  = "FEMALE"
)
