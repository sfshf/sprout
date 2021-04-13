package main

import (
	"context"
	"github.com/sfshf/sprout/app/govern/conf"
	"github.com/sfshf/sprout/repo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/url"
	"strings"
)

func NewMongoDB(ctx context.Context) (*mongo.Database, error) {
	c := conf.C.MongoDB
	srvUri, err := url.Parse(c.ServerUri)
	if err != nil {
		return nil, err
	}
	cliOpt := options.Client().SetHosts([]string{srvUri.Host})
	if direct := srvUri.Query().Get("directConnection"); direct != "" && strings.ToUpper(direct) == "TRUE" {
		cliOpt.SetDirect(true)
	}
	if dbName := srvUri.Path[1:]; dbName != "" {
		c.Database = dbName
	}
	client, err := mongo.NewClient(cliOpt)
	if err != nil {
		return nil, err
	}
	if err = client.Connect(ctx); err != nil {
		return nil, err
	}
	return client.Database(c.Database), nil
}

func InitRootAccount(ctx context.Context, repo *repo.Staff) error {
	c := conf.C.Root
	if sessionId, err := repo.UpsertRootAccount(ctx, c.Account, c.Password); err != nil {
		return err
	} else {
		c.SessionId = sessionId
		return nil
	}
}
