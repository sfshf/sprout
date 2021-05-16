package bll

import (
	"github.com/casbin/casbin/v2"
	b64Captcha "github.com/mojocn/base64Captcha"
	"github.com/sfshf/sprout/cache"
	"github.com/sfshf/sprout/pkg/jwtauth"
	"github.com/sfshf/sprout/repo"
)

type Menu struct {
	menuRepo *repo.Menu
}

func NewMenu(menuRepo *repo.Menu) *Menu {
	return &Menu{
		menuRepo: menuRepo,
	}
}

type Api struct {
	apiRepo   *repo.Api
	staffRepo *repo.Staff
}

func NewApi(apiRepo *repo.Api, staffRepo *repo.Staff) *Api {
	return &Api{
		apiRepo:   apiRepo,
		staffRepo: staffRepo,
	}
}

type Role struct {
	roleRepo  *repo.Role
	staffRepo *repo.Staff
	menuRepo  *repo.Menu
	apiRepo   *repo.Api
	enforcer  *casbin.Enforcer
}

func NewRole(roleRepo *repo.Role, staffRepo *repo.Staff, menuRepo *repo.Menu, apiRepo *repo.Api, enforcer *casbin.Enforcer) *Role {
	return &Role{
		roleRepo:  roleRepo,
		staffRepo: staffRepo,
		menuRepo:  menuRepo,
		apiRepo:   apiRepo,
		enforcer:  enforcer,
	}
}

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
