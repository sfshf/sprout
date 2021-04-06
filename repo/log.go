package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func AccessLogRepo() *AccessLog {
	return _accessLog
}

func InitAccessLogRepo(ctx context.Context, db *mongo.Database) {
	_accessLog = &AccessLog{
		coll: db.Collection(accessLogCollName),
	}
}

var (
	_accessLog *AccessLog

	accessLogCollName = "accessLog"
)

type AccessLog struct {
	coll *mongo.Collection
}

func (a *AccessLog) Collection() *mongo.Collection {
	return a.coll
}
