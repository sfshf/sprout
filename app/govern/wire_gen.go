// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/sfshf/sprout/app/govern/api"
	"github.com/sfshf/sprout/app/govern/bll"
	"github.com/sfshf/sprout/cache"
	"github.com/sfshf/sprout/repo"
)

import (
	_ "net/http/pprof"
)

// Injectors from wire.go:

func NewApp(ctx context.Context) (*App, func(), error) {
	database, cleanup, err := NewMongoDB(ctx)
	if err != nil {
		return nil, nil, err
	}
	accessLog := repo.NewAccessLogRepo(ctx, database)
	logger, err := NewLogger(accessLog)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	engine := NewRouter(ctx, logger)
	staff := repo.NewStaffRepo(ctx, database)
	jwtAuth := NewAuth()
	captcha := NewPictureCaptcha()
	bllStaff := bll.NewStaff(staff, jwtAuth, captcha)
	apiStaff := api.NewStaff(bllStaff)
	casbin := repo.NewCasbinRepo(ctx, database)
	enforcer := NewCasbin(ctx, casbin)
	bllCasbin := bll.NewCasbin(enforcer, staff)
	apiCasbin := api.NewCasbin(bllCasbin)
	bllAccessLog := bll.NewAccessLog(accessLog)
	apiAccessLog := api.NewAccessLog(bllAccessLog)
	user := repo.NewUserRepo(ctx, database)
	bllUser := bll.NewUser(user)
	apiUser := api.NewUser(bllUser)
	client, cleanup2, err := NewRedisDB(ctx)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	redisCache := cache.NewRedisCache(client)
	cacheCache, cleanup3, err := NewCache(ctx)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	memoryCache := cache.NewMemoryCache(cacheCache)
	app := &App{
		Router:        engine,
		StaffApi:      apiStaff,
		CasbinApi:     apiCasbin,
		AccessLogApi:  apiAccessLog,
		UserApi:       apiUser,
		StaffRepo:     staff,
		CasbinRepo:    casbin,
		UserRepo:      user,
		AccessLogRepo: accessLog,
		Auther:        jwtAuth,
		Enforcer:      enforcer,
		Redis:         redisCache,
		Cache:         memoryCache,
		PicCaptcha:    captcha,
		Logger:        logger,
	}
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

var (
	ApiSet   = wire.NewSet(api.NewStaff, api.NewCasbin, api.NewAccessLog, api.NewUser)
	BllSet   = wire.NewSet(bll.NewStaff, bll.NewCasbin, bll.NewAccessLog, bll.NewUser)
	RepoSet  = wire.NewSet(repo.NewStaffRepo, repo.NewCasbinRepo, repo.NewAccessLogRepo, repo.NewUserRepo)
	CacheSet = wire.NewSet(cache.NewMemoryCache, cache.NewRedisCache)
	AppSet   = wire.NewSet(
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
		NewRouter, wire.Struct(new(App), "*"),
	)
)
