package repo

import (
	"context"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewAccessLogRepo(ctx context.Context, db *mongo.Database) *AccessLog {
	a := &AccessLog{
		coll: db.Collection(accessLogCollName),
	}
	return a
}

var (
	accessLogCollName = "accessLog"
)

type AccessLog struct {
	coll *mongo.Collection
}

func (a *AccessLog) Collection() *mongo.Collection {
	return a.coll
}

func (a *AccessLog) CountByFilter(ctx context.Context, filter interface{}) (int64, error) {
	return a.coll.CountDocuments(ctx, filter, options.Count().SetMaxTime(time.Minute))
}

func (a *AccessLog) FindManyByFilter(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]model.AccessLog, error) {
	res := make([]model.AccessLog, 0)
	cursor, err := a.coll.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var one model.AccessLog
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
