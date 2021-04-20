package bll

import (
	"github.com/casbin/casbin/v2"
	b64Captcha "github.com/mojocn/base64Captcha"
	"github.com/sfshf/sprout/cache"
	"github.com/sfshf/sprout/pkg/jwtauth"
	"github.com/sfshf/sprout/repo"
)

type Staff struct {
	staffRepo  *repo.Staff
	redisCache *cache.RedisCache
	auther     *jwtauth.JWTAuth
	picCaptcha *b64Captcha.Captcha
}

func NewStaff(repo *repo.Staff, redisCache *cache.RedisCache, auther *jwtauth.JWTAuth, captcha *b64Captcha.Captcha) *Staff {
	return &Staff{
		staffRepo:  repo,
		redisCache: redisCache,
		auther:     auther,
		picCaptcha: captcha,
	}
}

type Casbin struct {
	enforcer  *casbin.Enforcer
	staffRepo *repo.Staff
}

func NewCasbin(enforcer *casbin.Enforcer, staffRepo *repo.Staff) *Casbin {
	return &Casbin{
		enforcer:  enforcer,
		staffRepo: staffRepo,
	}
}

type AccessLog struct {
	accessLog *repo.AccessLog
}

func NewAccessLog(repo *repo.AccessLog) *AccessLog {
	return &AccessLog{
		accessLog: repo,
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
