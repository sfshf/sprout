//+build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/sfshf/sprout/app/govern/api"
	"github.com/sfshf/sprout/app/govern/bll"
	"github.com/sfshf/sprout/repo"
)

var (
	ApiSet = wire.NewSet(
		api.NewStaff,
		api.NewCasbin,
		api.NewUser)
	BllSet = wire.NewSet(
		bll.NewStaff,
		bll.NewCasbin,
		bll.NewUser,
	)
	RepoSet = wire.NewSet(
		repo.NewStaffRepo,
		repo.NewCasbinRepo,
		repo.NewUserRepo,
		repo.NewAccessLogRepo,
	)
	AppSet = wire.NewSet(
		NewAuth,
		NewCasbin,
		NewPictureCaptcha,
		NewMongoDB,
		RepoSet,
		BllSet,
		ApiSet,
		NewRouter,
		wire.Struct(new(App), "*"),
	)
)

func NewApp(ctx context.Context) (*App, func(), error) {
	panic(wire.Build(AppSet))
}
