package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func UserRepo() *User {
	return _user
}

func InitUserRepo(ctx context.Context, db *mongo.Database) {
	_user = &User{
		coll: db.Collection(userCollName),
	}
	_user.coll.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{
				{"account", bsonx.Int32(1)},
			},
			Options: options.Index().SetUnique(true),
		},
	})
}

var (
	_user *User

	userCollName = "user"
)

type User struct {
	coll *mongo.Collection
}

func (a *User) Collection() *mongo.Collection {
	return a.coll
}
