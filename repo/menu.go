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

func NewMenuRepo(ctx context.Context, db *mongo.Database) *Menu {
	a := &Menu{
		coll: db.Collection(menuCollName),
	}
	a.coll.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{
				{"name", bsonx.Int32(1)},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{
				{"route", bsonx.Int32(1)},
			},
			Options: options.Index().SetUnique(true),
		},
	})
	return a
}

var (
	menuCollName = "menu"
)

type Menu struct {
	coll *mongo.Collection
}

func (a *Menu) Collection() *mongo.Collection {
	return a.coll
}

func (a *Menu) FindByID(ctx context.Context, id *primitive.ObjectID) (*model.Menu, error) {
	return nil, nil
}
