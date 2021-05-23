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

func NewStaffRepo(ctx context.Context, db *mongo.Database) *Staff {
	a := &Staff{
		coll: db.Collection(staffCollName),
	}
	a.coll.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{
				{"account", bsonx.Int32(1)},
			},
			Options: options.Index().SetUnique(true),
		},
	})
	return a
}

var (
	staffCollName = "staff"
)

type Staff struct {
	coll *mongo.Collection
}

func (a *Staff) Collection() *mongo.Collection {
	return a.coll
}

func (a *Staff) UpsertRootAccount(ctx context.Context, account, password string) (string, error) {
	var staff model.Staff
	if err := a.coll.FindOne(
		ctx,
		bson.M{
			"account": account,
			"role":    []string{model.RootRole},
		},
		options.FindOne().SetProjection(bson.D{{"_id", 1}}),
	).Decode(&staff); err != nil {
		if err != mongo.ErrNoDocuments {
			return "", err
		}
	} else {
		return staff.ID.Hex(), nil
	}
	salt := model.NewPasswdSalt()
	passwd := model.PasswdPtr(password, salt)
	res, err := a.coll.InsertOne(
		ctx,
		bson.M{
			"account":      account,
			"password":     passwd,
			"passwordSalt": salt,
			"role":         []string{model.RootRole},
			"signUpAt":     primitive.NewDateTimeFromTime(time.Now()),
		},
	)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (a *Staff) InsertOne(ctx context.Context, arg *model.Staff) error {
	res, err := a.coll.InsertOne(ctx, arg)
	if err != nil {
		return err
	}
	if res != nil {
		if id, is := res.InsertedID.(primitive.ObjectID); is {
			arg.ID = &id
		}
	}
	return nil
}

func (a *Staff) DeleteOne(ctx context.Context, argId *primitive.ObjectID) error {
	_, err := a.coll.DeleteOne(ctx, bson.M{"_id": argId})
	return err
}

func (a *Staff) FindOneByID(ctx context.Context, argId *primitive.ObjectID) (*model.Staff, error) {
	var m model.Staff
	if err := a.coll.FindOne(ctx, bson.M{"_id": argId}).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (a *Staff) FindOneByAccount(ctx context.Context, account string) (*model.Staff, error) {
	var m model.Staff
	if err := a.coll.FindOne(ctx, bson.M{"account": account}).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (a *Staff) TokenExists(ctx context.Context, id string, token string) bool {
	sessionId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	cnt, err := a.coll.CountDocuments(ctx, bson.M{"_id": sessionId, "signInToken": token}, options.Count().SetLimit(1))
	if err != nil || cnt < 1 {
		return false
	}
	return true
}

func (a *Staff) UpdateOneByID(ctx context.Context, arg *model.Staff) error {
	_, err := a.coll.UpdateOne(ctx, bson.M{"_id": arg.ID}, bson.M{"$set": arg})
	return err
}

func (a *Staff) FindManyByFilter(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]model.Staff, error) {
	res := make([]model.Staff, 0)
	cursor, err := a.coll.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var one model.Staff
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

func (a *Staff) CountByFilter(ctx context.Context, filter interface{}) (int64, error) {
	return a.coll.CountDocuments(ctx, filter, options.Count().SetMaxTime(time.Minute))
}

// https://docs.mongodb.com/manual/reference/operator/update/pull/#mongodb-update-up.-pull
func (a *Staff) EvictRole(ctx context.Context, argId *primitive.ObjectID, role *string) error {
	_, err := a.coll.UpdateOne(
		ctx,
		bson.D{
			{"_id", argId},
			{"roles", role},
		},
		bson.M{
			"$pull": bson.M{
				"roles": role,
			},
		})
	return err
}
