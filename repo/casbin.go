package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func CasbinRepo() *Casbin {
	return _casbin
}

func InitCasbinRepo(ctx context.Context, db *mongo.Database) {
	_casbin = &Casbin{
		coll: db.Collection(casbinCollName),
	}
	_casbin.coll.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{
				{"pType", bsonx.Int32(1)},
				{"v0", bsonx.Int32(1)},
				{"v1", bsonx.Int32(1)},
				{"v2", bsonx.Int32(1)},
				{"v3", bsonx.Int32(1)},
				{"v4", bsonx.Int32(1)},
			},
			Options: options.Index().SetUnique(true),
		},
	})
}

var (
	_casbin *Casbin

	casbinCollName = "casbin"
)

// A implementation of Adapter, BatchAdapter, FilteredAdapter interfaces of github.com/casbin/casbin/v2/persist package.
type Casbin struct {
	coll *mongo.Collection
}

func (a *Casbin) Collection() *mongo.Collection {
	return a.coll
}
