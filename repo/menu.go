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

func (a *Menu) FindOneByFilter(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (*model.Menu, error) {
	var res model.Menu
	if err := a.coll.FindOne(ctx, filter, opts...).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (a *Menu) FindOneByID(ctx context.Context, argId *primitive.ObjectID) (*model.Menu, error) {
	var res model.Menu
	if err := a.coll.FindOne(ctx, bson.M{"_id": argId}).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (a *Menu) FindOneAndDeleteByID(ctx context.Context, argId *primitive.ObjectID) (*model.Menu, error) {
	var res model.Menu
	if err := a.coll.FindOneAndDelete(ctx, bson.M{"_id": argId}).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (a *Menu) InsertOne(ctx context.Context, newM *model.Menu) error {
	res, err := a.coll.InsertOne(ctx, newM)
	if err != nil {
		return err
	}
	if res != nil {
		if id, is := res.InsertedID.(primitive.ObjectID); is {
			newM.ID = &id
		}
	}
	return nil
}

func (a *Menu) DeleteOneByID(ctx context.Context, argId *primitive.ObjectID) error {
	_, err := a.coll.DeleteOne(ctx, bson.M{"_id": argId})
	return err
}
