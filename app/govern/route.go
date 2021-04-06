package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/api"
	"github.com/sfshf/sprout/gin/middleware"
	"github.com/sfshf/sprout/pkg/jwtauth"
)

func NewController(auther *jwtauth.JWTAuth, enforcer *casbin.SyncedEnforcer, api *api.Api) *Controller {
	return &Controller{
		Auther:   auther,
		Enforcer: enforcer,
		Api:      api,
	}
}

type Controller struct {
	Auther   *jwtauth.JWTAuth
	Enforcer *casbin.SyncedEnforcer
	Api      *api.Api
}

func (a *Controller) InitRoutes(app *gin.Engine) {
	app.PUT("/signup", a.Api.Staff.SignUp)
	app.GET("/picCaptcha", a.Api.Staff.GetPicCaptcha)
	app.POST("/signin", a.Api.Staff.SignIn)

	v1 := app.Group("/api/v1")
	{
		v1.Use(middleware.JWT(a.Auther), middleware.Casbin(a.Enforcer))
		{
			v1.GET("/picCaptchaAnswer/:id", a.Api.Staff.GetPicCaptchaAnswer)
			v1.GET("/signout/:id", a.Api.Staff.SignOut)
			v1.DELETE("/signoff/:id", a.Api.Staff.SignOff)
		}

		casbin := v1.Group("/casbin")
		{
			policy := casbin.Group("/policy")
			{
				policy.PUT("", a.Api.Casbin.AddPolicy)
				policy.GET("/:id", a.Api.Casbin.Policy)
				policy.POST("/:id", a.Api.Casbin.UpdatePolicy)
				policy.DELETE("/:id", a.Api.Casbin.RemovePolicy)
				policy.GET("", a.Api.Casbin.Policies)
			}
		}

		staff := v1.Group("/staff")
		{
			staff.POST("/:id", a.Api.Staff.Update)
			staff.GET("/:id", a.Api.Staff.Info)
			staff.GET("", a.Api.Staff.List)
		}

		user := v1.Group("/user")
		{
			user.PUT("", a.Api.User.Add)
			user.DELETE("/:id", a.Api.User.Delete)
			user.POST("/:id", a.Api.User.Update)
			user.GET("/:id", a.Api.User.Info)
			user.GET("", a.Api.User.List)
		}
	}
}
