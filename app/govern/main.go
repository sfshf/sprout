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
	router, clear, err := NewGinEngine(ctx)
	if err != nil {
		panic(err)
	}
	defer clear()
	RunHTTPServer(ctx, router)

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
