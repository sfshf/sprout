package repo

import (
	"context"
	"github.com/sfshf/sprout/model"
	"github.com/sfshf/sprout/pkg/jwtauth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"time"
)

func StaffRepo() *Staff {
	return _staff
}

func InitStaffRepo(ctx context.Context, db *mongo.Database) {
	_staff = &Staff{
		coll: db.Collection(staffCollName),
	}
	_staff.coll.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{
				{"account", bsonx.Int32(1)},
			},
			Options: options.Index().SetUnique(true),
		},
	})
}

var (
	_staff *Staff

	staffCollName = "staff"
)

type Staff struct {
	coll *mongo.Collection
}

func (a *Staff) Collection() *mongo.Collection {
	return a.coll
}

func (a *Staff) UpsertRootAccount(ctx context.Context, account, password string) (string, error) {
	salt := model.NewPasswdSalt()
	passwd := model.PasswdPtr(password, salt)
	var staff model.Staff
	if err := a.coll.FindOne(
		ctx,
		bson.M{
			"account": account,
			"role":    model.RootRole,
		},
		options.FindOne().SetProjection(bson.D{{"_id", 1}}),
	).Decode(&staff); err != nil {
		if err != mongo.ErrNoDocuments {
			return "", err
		}
	} else {
		return staff.ID.Hex(), nil
	}
	res, err := a.coll.InsertOne(
		ctx,
		bson.M{
			"account":      account,
			"password":     passwd,
			"passwordSalt": salt,
			"role":         model.RootRole,
			"signUpAt":     primitive.NewDateTimeFromTime(time.Now()),
		},
	)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (a *Staff) InsertOne(ctx context.Context, m *model.Staff) error {
	res, err := a.coll.InsertOne(ctx, m)
	if err != nil {
		return err
	}
	if res != nil {
		if id, is := res.InsertedID.(primitive.ObjectID); is {
			m.ID = &id
		}
	}
	return nil
}

func (a *Staff) DeleteOne(ctx context.Context, id *primitive.ObjectID) error {
	_, err := a.coll.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (a *Staff) FindOneByAccount(ctx context.Context, username string) (*model.Staff, error) {
	var m model.Staff
	if err := a.coll.FindOne(ctx, bson.M{"account": username}).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (a *Staff) SignIn(ctx context.Context, id *primitive.ObjectID, token, ip *string, ts *primitive.DateTime) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"signInToken": token, "lastSignInIp": ip, "lastSignInTime": ts}}
	_, err := a.coll.UpdateOne(ctx, filter, update)
	return err
}

func (a *Staff) VerifySignInToken(ctx context.Context, id *primitive.ObjectID, token *string) error {
	cnt, err := a.coll.CountDocuments(ctx, bson.M{"_id": id, "signInToken": token}, options.Count().SetLimit(1))
	if err != nil || cnt < 1 {
		return jwtauth.ErrInvalidToken
	}
	return nil
}
