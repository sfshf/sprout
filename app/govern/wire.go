//+build wireinject

package main

import (
	"context"
	"github.com/gin-gonic/gin"
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
		repo.NewRoleRepo,
		repo.NewUserRepo,
		repo.NewAccessLogRepo,
	)
	ComponentSet = wire.NewSet(
		NewAuth,
		NewCasbin,
		NewPictureCaptcha,
		NewMongoDB,
		RepoSet,
		InitRootAccount,
		ApiSet,
		BllSet,
		wire.Struct(new(Controller), "*"),
		NewRouter,
	)
)

func NewGinEngine(ctx context.Context) (*gin.Engine, func(), error) {
	panic(wire.Build(ComponentSet))
}
