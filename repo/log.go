package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
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
