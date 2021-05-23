package repo

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"time"
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

func (a *Api) FindOneByID(ctx context.Context, argId *primitive.ObjectID) (*model.Api, error) {
	var m model.Api
	if err := a.coll.FindOne(ctx, bson.M{"_id": argId}).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (a *Api) FindOneByFilter(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (*model.Api, error) {
	var m model.Api
	if err := a.coll.FindOne(ctx, filter, opts...).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (a *Api) FindOneAndDeleteByID(ctx context.Context, argId *primitive.ObjectID) (*model.Api, error) {
	var m model.Api
	if err := a.coll.FindOneAndDelete(ctx, bson.M{"_id": argId}).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (a *Api) FindOneAndUpdateByID(ctx context.Context, arg *model.Api) (*model.Api, error) {
	var m model.Api
	if err := a.coll.FindOneAndUpdate(ctx, bson.M{"_id": arg.ID}, bson.M{"$set": arg}).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (a *Api) InsertOne(ctx context.Context, newM *model.Api) error {
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

func (a *Api) DeleteOneByID(ctx context.Context, argId *primitive.ObjectID) error {
	_, err := a.coll.DeleteOne(ctx, bson.M{"_id": argId})
	return err
}

func (a *Api) UpdateOneByID(ctx context.Context, arg *model.Api) error {
	_, err := a.coll.UpdateOne(ctx, bson.M{"_id": arg.ID}, bson.M{"$set": arg})
	return err
}

func (a *Api) CountByFilter(ctx context.Context, filter interface{}) (int64, error) {
	return a.coll.CountDocuments(ctx, filter, options.Count().SetMaxTime(time.Minute))
}

func (a *Api) FindManyByFilter(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]model.Api, error) {
	res := make([]model.Api, 0)
	cursor, err := a.coll.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var one model.Api
		if err := cursor.Decode(&one); err != nil {
			return nil, err
		}
		res = append(res, one)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return res, nil
}
