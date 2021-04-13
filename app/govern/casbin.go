package main

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/sfshf/sprout/app/govern/conf"
	"github.com/sfshf/sprout/repo"
	"time"
)

func NewCasbin(ctx context.Context, repo *repo.Casbin) (*casbin.SyncedEnforcer, func()) {
	c := conf.C.Casbin
	if c.Model == "" {
		c.Model = "app/govern/conf/casbin_rbac.model"
	}
	e, err := casbin.NewSyncedEnforcer(c.Model)
	if err != nil {
		panic(err)
	}
	e.EnableLog(c.Debug)
	err = e.InitWithModelAndAdapter(e.GetModel(), repo)
	if err != nil {
		panic(err)
	}
	e.EnableEnforce(c.Enable)

	deferFunc := func() {}
	if c.AutoLoad {
		e.StartAutoLoadPolicy(time.Duration(c.AutoLoadInternal) * time.Second)
		deferFunc = func() {
			e.StopAutoLoadPolicy()
		}
	}
	return e, deferFunc
}
