package repo

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func NewApiRepo(ctx context.Context, db *mongo.Database) *Api {
	a := &Api{
		coll: db.Collection(apiCollName),
	}
	a.coll.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{
				{"method", bsonx.Int32(1)},
				{"path", bsonx.Int32(1)},
			},
			Options: options.Index().SetUnique(true),
		},
	})
	return a
}

var (
	apiCollName = "api"
)

type Api struct {
	coll *mongo.Collection
}

func (a *Api) Collection() *mongo.Collection {
	return a.coll
}

func (a *Api) FindByID(ctx context.Context, id *primitive.ObjectID) (*model.Api, error) {
	return nil, nil
}
