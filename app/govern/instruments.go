package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	b64Captcha "github.com/mojocn/base64Captcha"
	"github.com/sfshf/sprout/app/govern/config"
	"github.com/sfshf/sprout/app/govern/ginx"
	"github.com/sfshf/sprout/pkg/jwtauth"
	"github.com/sfshf/sprout/pkg/logger"
	"github.com/sfshf/sprout/repo"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Logger --------------------------------------------------------------------------------------------------------------

func NewLogger(repo *repo.AccessLog) (*logger.Logger, error) {
	c := config.C.Log
	var writers []io.Writer
	if !c.SkipStdout {
		writers = append(writers, os.Stderr)
	}
	if c.Log2Mongo {
		writer, err := logger.MongoWriter(repo.Collection())
		if err != nil {
			return nil, err
		}
		writers = append(writers, writer)
	}
	log.Println("Access logger is on!!!")
	return logger.NewLogger(writers...), nil
}

// Router --------------------------------------------------------------------------------------------------------------

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
		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	log.Println("App router is on!!!")
	return app
}

// Authenticator -------------------------------------------------------------------------------------------------------

func NewAuth() *jwtauth.JWTAuth {
	c := config.C.JWTAuth
	var opts []jwtauth.Option
	if c.Expired > 0 {
		opts = append(opts, jwtauth.SetExpired(c.Expired))
	}
	if c.SigningKey != "" {
		opts = append(opts, jwtauth.SetSigningKey([]byte(c.SigningKey)))
		opts = append(opts, jwtauth.SetKeyFunc(func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwtauth.ErrInvalidToken
			}
			return []byte(c.SigningKey), nil
		}))
	}
	log.Println("JWT auther is on!!!")
	return jwtauth.New(opts...)
}

// Authorizer ----------------------------------------------------------------------------------------------------------

func NewCasbin(ctx context.Context, repo *repo.Casbin) *casbin.Enforcer {
	c := config.C.Casbin
	if c.Model == "" {
		c.Model = "app/govern/config/casbin_rbac.model"
	}
	enforcer, err := casbin.NewEnforcer(c.Model)
	if err != nil {
		panic(err)
	}
	enforcer.EnableLog(c.Debug)
	err = enforcer.InitWithModelAndAdapter(enforcer.GetModel(), repo)
	if err != nil {
		panic(err)
	}
	enforcer.EnableEnforce(c.Enable)
	log.Println("Casbin enforcer is on!!!")
	return enforcer
}

// Captcha -------------------------------------------------------------------------------------------------------------

func NewPictureCaptcha() *b64Captcha.Captcha {
	c := config.C.PicCaptcha
	driver := b64Captcha.NewDriverDigit(c.Height, c.Width, c.Length, c.MaxSkew, c.DotCount)
	var store b64Captcha.Store
	if c.RedisStore {
		// TODO Redis store.
	} else {
		store = b64Captcha.NewMemoryStore(c.Threshold, c.Expiration)
	}
	log.Println("Picture captcha is on!!!")
	return b64Captcha.NewCaptcha(driver, store)
}
