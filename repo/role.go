package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func NewRoleRepo(ctx context.Context, db *mongo.Database) *Role {
	a := &Role{
		coll: db.Collection(roleCollName),
	}
	a.coll.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{
				{"name", bsonx.Int32(1)},
			},
			Options: options.Index().SetUnique(true),
		},
	})
	return a
}

var (
	roleCollName = "role"
)

type Role struct {
	coll *mongo.Collection
}

func (a *Role) Collection() *mongo.Collection {
	return a.coll
}
