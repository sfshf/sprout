package main

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/sfshf/sprout/app/govern/config"
	"github.com/sfshf/sprout/repo"
)

func NewCasbin(ctx context.Context, repo *repo.Casbin) *casbin.Enforcer {
	c := config.C.Casbin
	if c.Model == "" {
		c.Model = "app/govern/config/casbin_rbac.model"
	}
	enforcer, err := casbin.NewEnforcer(c.Model)
	if err != nil {
		panic(err)
	}
	enforcer.EnableLog(c.Debug)
	err = enforcer.InitWithModelAndAdapter(enforcer.GetModel(), repo)
	if err != nil {
		panic(err)
	}
	enforcer.EnableEnforce(c.Enable)
	return enforcer
}
