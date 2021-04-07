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

func InitRepos(ctx context.Context) {
	c := conf.C.MongoDB
	srvUri, err := url.Parse(c.ServerUri)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	if err = client.Connect(ctx); err != nil {
		panic(err)
	}
	db := client.Database(c.Database)
	repo.InitStaffRepo(ctx, db)
	repo.InitRoleRepo(ctx, db)
	repo.InitCasbinRepo(ctx, db)
	repo.InitUserRepo(ctx, db)
	repo.InitAccessLogRepo(ctx, db)
}

func InitRootAccount(ctx context.Context) {
	c := conf.C.Root
	staffRepo := repo.StaffRepo()
	if sessionId, err := staffRepo.UpsertRootAccount(ctx, c.Account, c.Password); err != nil {
		panic(err)
	} else {
		c.SessionId = sessionId
	}
}
