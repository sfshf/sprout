package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	b64Captcha "github.com/mojocn/base64Captcha"
	"github.com/sfshf/sprout/app/govern/api"
	"github.com/sfshf/sprout/app/govern/config"
	"github.com/sfshf/sprout/app/govern/ginx"
	"github.com/sfshf/sprout/cache"
	"github.com/sfshf/sprout/pkg/jwtauth"
	"github.com/sfshf/sprout/pkg/logger"
	"github.com/sfshf/sprout/repo"
)

type App struct {
	StaffApi     *api.Staff
	RoleApi      *api.Role
	MenuApi      *api.Menu
	ApiApi       *api.Api
	CasbinApi    *api.Casbin
	AccessLogApi *api.AccessLog
	UserApi      *api.User

	StaffRepo     *repo.Staff
	CasbinRepo    *repo.Casbin
	UserRepo      *repo.User
	AccessLogRepo *repo.AccessLog
	RedisCache    *cache.RedisCache
	MemoryCache   *cache.MemoryCache

	Router     *gin.Engine
	Auther     *jwtauth.JWTAuth
	Enforcer   *casbin.Enforcer
	PicCaptcha *b64Captcha.Captcha
	Logger     *logger.Logger
}

func (a *App) InitRootAccount(ctx context.Context) error {
	c := config.C.Root
	if sessionId, err := a.StaffRepo.UpsertRootAccount(ctx, c.Account, c.Password); err != nil {
		return err
	} else {
		c.SessionId = sessionId
		return nil
	}
}

func (a *App) RunHTTPServer(ctx context.Context) func() {
	c := config.C.HTTP
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      a.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	go func() {
		log.Printf("HTTP server is running at %s", addr)
		var err error
		if c.CertFile != "" && c.CertKeyFile != "" {
			srv.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
			err = srv.ListenAndServeTLS(c.CertFile, c.CertKeyFile)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(c.ShutdownTimeout))
		defer cancel()
		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			log.Println(err.Error())
		}
	}
}

func (a *App) InitRoutes(ctx context.Context) {
	v1 := a.Router.Group("/api/v1")
	{
		v1.POST("/signUp", a.StaffApi.SignUp)
		v1.GET("/picCaptcha", a.StaffApi.GetPicCaptcha)
		v1.PATCH("/signIn", a.StaffApi.SignIn)

		v1.Use(ginx.JWT(a.Auther, a.RedisCache))

		{
			v1.GET("/picCaptchaAnswer", a.StaffApi.GetPicCaptchaAnswer)
			v1.PATCH("/signOut", a.StaffApi.SignOut)
			v1.DELETE("/signOff/:id", a.StaffApi.SignOff)
		}

		v1.Use(ginx.Casbin(a.Enforcer, config.C.Root.SessionId))

		staff := v1.Group("/staff")
		{
			staff.PATCH("/:id/password", a.StaffApi.UpdatePassword)
			staff.PATCH("/:id/email", a.StaffApi.UpdateEmail)
			staff.PATCH("/:id/phone", a.StaffApi.UpdatePhone)
			staff.PATCH("/:id/roles", ginx.MustRoot(config.C.Root.SessionId), a.StaffApi.UpdateRoles)
			staff.PATCH("/:id/signInIpWhitelist", a.StaffApi.UpdateSignInIpWhitelist)
			staff.GET("/:id", a.StaffApi.Profile)
			staff.GET("", a.StaffApi.List)
			staff.PATCH("/:id/enable", ginx.MustRoot(config.C.Root.SessionId), a.StaffApi.Enable)
		}

		role := v1.Group("/role", ginx.MustRoot(config.C.Root.SessionId))
		{
			role.POST("", a.RoleApi.Add)
			role.DELETE("/:id", a.RoleApi.Evict)
			role.PUT("/:id", a.RoleApi.Update)
			role.PUT("/:id/authorize", a.RoleApi.Authorize)
			role.GET("/:id", a.RoleApi.Profile)
			role.GET("", a.RoleApi.List)
			role.PATCH("/:id/enable", a.RoleApi.Enable)
		}

		menu := v1.Group("/menu", ginx.MustRoot(config.C.Root.SessionId))
		{
			menu.POST("", a.MenuApi.Add)
			menu.DELETE("/:id", a.MenuApi.Evict)
			menu.PUT("/:id", a.MenuApi.Update)
			menu.GET("/:id", a.MenuApi.Profile)
			menu.GET("", a.MenuApi.List)
			menu.PATCH("/:id/enable", a.MenuApi.Enable)
		}

		api := v1.Group("/api", ginx.MustRoot(config.C.Root.SessionId))
		{
			api.POST("", a.ApiApi.Add)
			api.DELETE("/:id", a.ApiApi.Evict)
			api.PUT("/:id", a.ApiApi.Update)
			api.GET("/:id", a.ApiApi.Profile)
			api.GET("", a.ApiApi.List)
			api.PATCH("/:id/enable", a.ApiApi.Enable)
		}

		casbin := v1.Group("/casbin", ginx.MustRoot(config.C.Root.SessionId))
		{
			priority := casbin.Group("/priority")
			{
				priority.GET("", a.CasbinApi.Priorities)
			}

			object := casbin.Group("/resource")
			{
				object.GET("", a.CasbinApi.AllApiObjActMap)
			}

			policy := casbin.Group("/policy")
			{
				//policy.POST("", a.CasbinApi.AddPolicy)
				policy.GET("/:role", a.CasbinApi.PoliciesOfRole)
				//policy.PUT("/:id", a.CasbinApi.UpdatePolicy)
				//policy.DELETE("/:id", a.CasbinApi.RemovePolicy)
				//policy.GET("", a.CasbinApi.Policies)
			}

			role := casbin.Group("/role")
			{
				role.POST("", a.CasbinApi.AddRole)
				role.DELETE("/:role", a.CasbinApi.DeleteRole)
				role.PUT("/:role/set/:staffId", a.CasbinApi.SetRole)
				role.DELETE("/:role/unset/:staffId", a.CasbinApi.UnsetRole)
				role.GET("", a.CasbinApi.AllRoles)
				role.GET("/:staffId", a.CasbinApi.RolesOfStaff)
			}

			staff := casbin.Group("/staff")
			{
				staff.GET("/:role", a.CasbinApi.StaffsOfRole)
			}
		}

		accessLog := v1.Group("/accessLog")
		{
			accessLog.GET("", a.AccessLogApi.List)
		}

		user := v1.Group("/user")
		{
			user.POST("", a.UserApi.Add)
			user.DELETE("/:id", a.UserApi.Delete)
			user.PUT("/:id", a.UserApi.Update)
			user.GET("/:id", a.UserApi.Info)
			user.GET("", a.UserApi.List)
		}
	}
	a.CasbinApi.Routes = a.Router.Routes()
}
