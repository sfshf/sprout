package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func RoleRepo() *Role {
	return _role
}

func InitRoleRepo(ctx context.Context, db *mongo.Database) {
	_role = &Role{
		coll: db.Collection(roleCollName),
	}
	_role.coll.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{
				{"name", bsonx.Int32(1)},
			},
			Options: options.Index().SetUnique(true),
		},
	})
}

var (
	_role *Role

	roleCollName = "role"
)

type Role struct {
	coll *mongo.Collection
}

func (a *Role) Collection() *mongo.Collection {
	return a.coll
}
