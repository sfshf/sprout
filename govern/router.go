package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/govern/conf"
	"github.com/sfshf/sprout/govern/internal/ginx/middleware"
	swag "github.com/swaggo/gin-swagger"
	swagFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
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
	return app
}
