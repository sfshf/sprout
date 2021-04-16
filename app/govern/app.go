package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	b64Captcha "github.com/mojocn/base64Captcha"
	"github.com/sfshf/sprout/app/govern/api"
	"github.com/sfshf/sprout/app/govern/config"
	"github.com/sfshf/sprout/app/govern/ginx/middleware"
	"github.com/sfshf/sprout/pkg/jwtauth"
	"github.com/sfshf/sprout/repo"
	"log"
	"net/http"
	"time"
)

type App struct {
	Router    *gin.Engine
	StaffApi  *api.Staff
	CasbinApi *api.Casbin
	UserApi   *api.User

	StaffRepo     *repo.Staff
	CasbinRepo    *repo.Casbin
	UserRepo      *repo.User
	AccessLogRepo *repo.AccessLog

	Auther     *jwtauth.JWTAuth
	Enforcer   *casbin.SyncedEnforcer
	PicCaptcha *b64Captcha.Captcha
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
		v1.PUT("/signUp", a.StaffApi.SignUp)
		v1.GET("/picCaptcha", a.StaffApi.GetPicCaptcha)
		v1.POST("/signIn", a.StaffApi.SignIn)

		v1.Use(middleware.JWT(a.Auther, a.StaffRepo))

		{
			v1.GET("/picCaptchaAnswer/:id", a.StaffApi.GetPicCaptchaAnswer)
			v1.GET("/signOut", a.StaffApi.SignOut)
			v1.DELETE("/signOff/:id", a.StaffApi.SignOff)
		}

		v1.Use(middleware.Casbin(a.Enforcer, config.C.Root.SessionId))

		staff := v1.Group("/staff")
		{
			staff.POST("/:id", a.StaffApi.Update)
			staff.GET("/:id", a.StaffApi.Profile)
			staff.GET("", a.StaffApi.List)
		}

		casbin := v1.Group("/casbin")
		{
			policy := casbin.Group("/policy")
			{
				policy.PUT("", a.CasbinApi.AddPolicy)
				policy.GET("/:id", a.CasbinApi.Policy)
				policy.POST("/:id", a.CasbinApi.UpdatePolicy)
				policy.DELETE("/:id", a.CasbinApi.RemovePolicy)
				policy.GET("", a.CasbinApi.Policies)
			}
		}

		user := v1.Group("/user")
		{
			user.PUT("", a.UserApi.Add)
			user.DELETE("/:id", a.UserApi.Delete)
			user.POST("/:id", a.UserApi.Update)
			user.GET("/:id", a.UserApi.Info)
			user.GET("", a.UserApi.List)
		}
	}
}
