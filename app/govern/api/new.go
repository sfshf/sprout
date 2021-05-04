package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/bll"
)

type Menu struct {
	bll *bll.Menu
}

func NewMenu(bll *bll.Menu) *Menu {
	return &Menu{
		bll: bll,
	}
}

type Api struct {
	bll *bll.Api
}

func NewApi(bll *bll.Api) *Api {
	return &Api{
		bll: bll,
	}
}

type Role struct {
	bll *bll.Role
}

func NewRole(bll *bll.Role) *Role {
	return &Role{
		bll: bll,
	}
}

type Staff struct {
	bll *bll.Staff
}

func NewStaff(bll *bll.Staff) *Staff {
	return &Staff{
		bll: bll,
	}
}

type Casbin struct {
	bll    *bll.Casbin
	Routes gin.RoutesInfo
}

func NewCasbin(bll *bll.Casbin) *Casbin {
	return &Casbin{
		bll: bll,
	}
}

type AccessLog struct {
	bll *bll.AccessLog
}

func NewAccessLog(bll *bll.AccessLog) *AccessLog {
	return &AccessLog{
		bll: bll,
	}
}

type User struct {
	bll *bll.User
}

func NewUser(bll *bll.User) *User {
	return &User{
		bll: bll,
	}
}
