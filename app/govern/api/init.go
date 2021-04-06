package api

import (
	b64Captcha "github.com/mojocn/base64Captcha"
	"github.com/sfshf/sprout/app/govern/bll"
	"github.com/sfshf/sprout/pkg/jwtauth"
)

func Init(auther *jwtauth.JWTAuth, captcha *b64Captcha.Captcha) *Api {
	bll := bll.Init(auther, captcha)
	return &Api{
		Casbin: &Casbin{
			bll: bll.Casbin,
		},
		Staff: &Staff{
			bll: bll.Staff,
		},
		User: &User{
			bll: bll.User,
		},
	}
}

type Api struct {
	Casbin *Casbin
	Staff  *Staff
	User   *User
}

type Casbin struct {
	bll *bll.Casbin
}

type Staff struct {
	bll *bll.Staff
}

type User struct {
	bll *bll.User
}
