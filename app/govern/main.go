package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/sfshf/sprout/app/govern/docs"
)

// @title Govern APIs
// @version 0.0.1-beta
// @description This is a back-end government app.
// @termsOfService http://swagger.io/terms/

// @contact.name gavin
// @contact.url http://github.com/sfshf
// @contact.email sfshf@github.com

// @license.name MIT
// @license.url https://github.com/sfshf/sprout/blob/main/LICENSE

// @host localhost:8000
// @basePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx := context.Background()
	app, clear, err := NewApp(ctx)
	if err != nil {
		panic(err)
	}
	defer clear()
	app.InitRootAccount(ctx)
	app.InitRoutes(ctx)
	app.RunHTTPServer(ctx)

	go func() {
		log.Println(http.ListenAndServe(":8010", nil))
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
