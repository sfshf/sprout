package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/bll"
)

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
