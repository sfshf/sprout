package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID            *primitive.ObjectID `bson:"_id,omitempty"`
	Account       *string             `bson:"account,omitempty"`
	Password      *string             `bson:"password,omitempty"`
	PasswordSalt  *string             `bson:"passwordSalt,omitempty"`
	Nickname      *string             `bson:"nickname,omitempty"`
	RealName      *string             `bson:"realName,omitempty"`
	Avatar        *string             `bson:"avatar,omitempty"`
	QRCode        *string             `bson:"qrCode,omitempty"`
	MoreInfo      *MoreInfo           `bson:"moreInfo,omitempty"`
	LastLoginTime *primitive.DateTime `bson:"lastLoginTime,omitempty"`
	CreatedTime   *primitive.DateTime `bson:"createdTime,omitempty"`
	DeletedTime   *primitive.DateTime `bson:"deletedTime,omitempty"`
}

type MoreInfo struct {
	Gender  string `bson:"gender,omitempty"`
	Region  string `bson:"region,omitempty"`
	WhatsUp string `bson:"whatsUp,omitempty"`
}
