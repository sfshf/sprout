package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/api"
	"github.com/sfshf/sprout/app/govern/conf"
	"github.com/sfshf/sprout/gin/middleware"
	swag "github.com/swaggo/gin-swagger"
	swagFiles "github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title
// @version
// @description
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http https
// @basePath /
// @contact.name sfshf
func main() {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx := context.Background()
	conf.Init()
	auther := InitAuth()
	captcha := InitPictureCaptcha()
	InitRepos(ctx)
	InitRootAccount(ctx)
	enforcer, deferFunc := InitCasbin(ctx)
	defer deferFunc()
	api := api.Init(auther, captcha)
	controller := NewController(auther, enforcer, api)
	router := InitGinEngine()
	controller.InitRoutes(router)
	InitHTTPServer(ctx, router)

	go func() {
		log.Println(http.ListenAndServe(":8090", nil))
	}()

EXIT:
	for {
		sig := <-sc
		log.Printf("Signal[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}
	log.Println("Server Exit")
	time.Sleep(time.Second)
	os.Exit(state)
}

func InitGinEngine() *gin.Engine {
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
	if dir := conf.C.WWW; dir != "" {
		app.Use(middleware.WWW(dir))
	}
	return app
}
