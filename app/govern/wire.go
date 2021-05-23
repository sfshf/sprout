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
		api.NewRole,
		api.NewMenu,
		api.NewApi,
		api.NewCasbin,
		api.NewAccessLog,
	)
	BllSet = wire.NewSet(
		bll.NewStaff,
		bll.NewRole,
		bll.NewMenu,
		bll.NewApi,
		bll.NewCasbin,
		bll.NewAccessLog,
	)
	RepoSet = wire.NewSet(
		repo.NewStaffRepo,
		repo.NewRoleRepo,
		repo.NewMenuRepo,
		repo.NewApiRepo,
		repo.NewCasbinRepo,
		repo.NewAccessLogRepo,
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
