package bll

import (
	b64Captcha "github.com/mojocn/base64Captcha"
	"github.com/sfshf/sprout/govern/internal/pkg/jwtauth"
	"github.com/sfshf/sprout/repo"
)

type Staff struct {
	staffRepo  *repo.Staff
	auther     *jwtauth.JWTAuth
	picCaptcha *b64Captcha.Captcha
}

func NewStaff(repo *repo.Staff, auther *jwtauth.JWTAuth, captcha *b64Captcha.Captcha) *Staff {
	return &Staff{
		staffRepo:  repo,
		auther:     auther,
		picCaptcha: captcha,
	}
}

type Casbin struct {
	casbinRepo *repo.Casbin
}

func NewCasbin(repo *repo.Casbin) *Casbin {
	return &Casbin{
		casbinRepo: repo,
	}
}

type User struct {
	userRepo *repo.User
}

func NewUser(repo *repo.User) *User {
	return &User{
		userRepo: repo,
	}
}
