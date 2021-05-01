package ginx

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/config"
)

// TODO:
func CORS() gin.HandlerFunc {
	c := config.C.CORS
	return cors.New(
		cors.Config{
			AllowAllOrigins: true,
			//AllowOrigins:    c.AllowOrigins,
			//AllowOriginFunc:        nil,
			AllowMethods:     c.AllowMethods,
			AllowHeaders:     c.AllowHeaders,
			AllowCredentials: c.AllowCredentials,
			//ExposeHeaders:          nil,
			MaxAge: c.MaxAge,
			//AllowWildcard:          false,
			//AllowBrowserExtensions: false,
			//AllowWebSockets:        false,
			//AllowFiles:             false,
		},
	)
}
