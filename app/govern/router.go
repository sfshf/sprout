package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/api"
	"github.com/sfshf/sprout/app/govern/conf"
	"github.com/sfshf/sprout/gin/middleware"
	"github.com/sfshf/sprout/pkg/jwtauth"
	"github.com/sfshf/sprout/repo"
	swag "github.com/swaggo/gin-swagger"
	swagFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter(ctrl *Controller) *gin.Engine {
	gin.SetMode(conf.C.RunMode)

	app := gin.New()
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())
	// TODO Custom access logger
	app.Use(gin.Logger())
	// TODO Custom recovery logger
	app.Use(gin.Recovery())
	// TODO CORS middleware
	// TODO TraceID middleware
	// TODO GZIP
	if conf.C.Swagger {
		app.GET("/swagger/*any", swag.WrapHandler(swagFiles.Handler))
	}
	ctrl.InitRoutes(app)
	return app
}

type Controller struct {
	Auther    *jwtauth.JWTAuth
	Enforcer  *casbin.SyncedEnforcer
	Staff     *api.Staff
	Casbin    *api.Casbin
	User      *api.User
	StaffRepo *repo.Staff
}

func (a *Controller) InitRoutes(app *gin.Engine) {
	app.PUT("/signUp", a.Staff.SignUp)
	app.GET("/picCaptcha", a.Staff.GetPicCaptcha)
	app.POST("/signIn", a.Staff.SignIn)
	app.Use(middleware.JWT(a.Auther, a.StaffRepo), middleware.Casbin(a.Enforcer, conf.C.Root.SessionId))
	{
		app.GET("/picCaptchaAnswer/:id", a.Staff.GetPicCaptchaAnswer)
		app.GET("/signOut", a.Staff.SignOut)
		app.DELETE("/signOff/:id", a.Staff.SignOff)
	}

	v1 := app.Group("/api/v1")
	{
		casbin := v1.Group("/casbin")
		{
			policy := casbin.Group("/policy")
			{
				policy.PUT("", a.Casbin.AddPolicy)
				policy.GET("/:id", a.Casbin.Policy)
				policy.POST("/:id", a.Casbin.UpdatePolicy)
				policy.DELETE("/:id", a.Casbin.RemovePolicy)
				policy.GET("", a.Casbin.Policies)
			}
		}

		staff := v1.Group("/staff")
		{
			staff.POST("/:id", a.Staff.Update)
			staff.GET("/:id", a.Staff.Profile)
			staff.GET("", a.Staff.List)
		}

		user := v1.Group("/user")
		{
			user.PUT("", a.User.Add)
			user.DELETE("/:id", a.User.Delete)
			user.POST("/:id", a.User.Update)
			user.GET("/:id", a.User.Info)
			user.GET("", a.User.List)
		}
	}
}
