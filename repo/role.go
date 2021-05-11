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

func (a *Role) InsertOne(ctx context.Context, newM *model.Role) error {
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

func (a *Role) FindOneByID(ctx context.Context, argId *primitive.ObjectID) (*model.Role, error) {
	var res model.Role
	if err := a.coll.FindOne(ctx, bson.M{"_id": argId}).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (a *Role) EvictRole(ctx context.Context, argId *primitive.ObjectID) error {
	_, err := a.coll.DeleteOne(ctx, bson.M{"_id": argId})
	return err
}

func (a *Role) UpdateOneByID(ctx context.Context, arg *model.Role) error {
	_, err := a.coll.UpdateOne(ctx, bson.M{"_id": arg.ID}, bson.M{"$set": arg})
	return err
}
