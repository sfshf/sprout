//+build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/sfshf/sprout/app/govern/api"
	"github.com/sfshf/sprout/app/govern/bll"
	"github.com/sfshf/sprout/cache"
	"github.com/sfshf/sprout/repo"
)

var (
	ApiSet = wire.NewSet(
		api.NewStaff,
		api.NewCasbin,
		api.NewAccessLog,
		api.NewUser,
	)
	BllSet = wire.NewSet(
		bll.NewStaff,
		bll.NewCasbin,
		bll.NewAccessLog,
		bll.NewUser,
	)
	RepoSet = wire.NewSet(
		repo.NewStaffRepo,
		repo.NewCasbinRepo,
		repo.NewAccessLogRepo,
		repo.NewUserRepo,
	)
	CacheSet = wire.NewSet(
		cache.NewMemoryCache,
		cache.NewRedisCache,
	)
	AppSet = wire.NewSet(
		NewAuth,
		NewCasbin,
		NewPictureCaptcha,
		NewMongoDB,
		NewRedisDB,
		NewCache,
		NewLogger,
		RepoSet,
		CacheSet,
		BllSet,
		ApiSet,
		NewRouter,
		wire.Struct(new(App), "*"),
	)
)

func NewApp(ctx context.Context) (*App, func(), error) {
	panic(wire.Build(AppSet))
}
