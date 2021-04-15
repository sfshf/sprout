//+build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/sfshf/sprout/govern/api"
	"github.com/sfshf/sprout/govern/bll"
	"github.com/sfshf/sprout/govern/ginx/router"
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
		repo.NewRoleRepo,
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
		router.NewRouter,
		wire.Struct(new(App), "*"),
	)
)

func NewApp(ctx context.Context) (*App, func(), error) {
	panic(wire.Build(AppSet))
}
