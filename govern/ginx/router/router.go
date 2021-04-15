package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/govern/config"
	"github.com/sfshf/sprout/govern/ginx/middleware"
	swag "github.com/swaggo/gin-swagger"
	swagFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	gin.SetMode(config.C.RunMode)

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
	if config.C.Swagger {
		app.GET("/swagger/*any", swag.WrapHandler(swagFiles.Handler))
	}
	return app
}
