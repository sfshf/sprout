package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/config"
	"github.com/sfshf/sprout/app/govern/ginx"
	"github.com/sfshf/sprout/pkg/logger"
	swag "github.com/swaggo/gin-swagger"
	swagFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter(ctx context.Context, logger *logger.Logger) *gin.Engine {
	gin.SetMode(config.C.RunMode)

	app := gin.New()
	app.NoMethod(ginx.NoMethodHandler())
	app.NoRoute(ginx.NoRouteHandler())
	app.Use(ginx.Logger(logger, config.C.Log.Enable))
	// TODO Custom recovery logger
	app.Use(gin.Recovery())
	app.Use(ginx.CORS())
	app.Use(ginx.TraceId())
	app.Use(ginx.GZIP())
	if config.C.Swagger {
		app.GET("/swagger/*any", swag.WrapHandler(swagFiles.Handler))
	}
	return app
}
