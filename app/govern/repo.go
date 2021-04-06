package main

import (
	"context"
	"github.com/sfshf/sprout/app/govern/conf"
	"github.com/sfshf/sprout/repo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitRepos(ctx context.Context) {
	c := conf.C.MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(c.ClientUri))
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
