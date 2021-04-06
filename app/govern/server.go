package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/sfshf/sprout/app/govern/conf"
	"log"
	"net/http"
	"time"
)

func InitHTTPServer(ctx context.Context, handler http.Handler) func() {
	c := conf.C.HTTP
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	go func() {
		log.Printf("HTTP server is running at %s", addr)
		var err error
		if c.CertFile != "" && c.KeyFile != "" {
			srv.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
			err = srv.ListenAndServeTLS(c.CertFile, c.KeyFile)
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
