package bll

import (
	b64Captcha "github.com/mojocn/base64Captcha"
	"github.com/sfshf/sprout/pkg/jwtauth"
	"github.com/sfshf/sprout/repo"
)

func Init(auther *jwtauth.JWTAuth, captcha *b64Captcha.Captcha) *Bll {
	return &Bll{
		Casbin: &Casbin{
			casbinRepo: repo.CasbinRepo(),
		},
		Staff: &Staff{
			staffRepo:  repo.StaffRepo(),
			auther:     auther,
			picCaptcha: captcha,
		},
		User: &User{
			userRepo: repo.UserRepo(),
		},
	}
}

type Bll struct {
	Casbin *Casbin
	Staff  *Staff
	User   *User
}

type Casbin struct {
	casbinRepo *repo.Casbin
}

type Staff struct {
	staffRepo  *repo.Staff
	auther     *jwtauth.JWTAuth
	picCaptcha *b64Captcha.Captcha
}

type User struct {
	userRepo *repo.User
}
