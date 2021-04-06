package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/schema"
	"net/http"
)

func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(
			http.StatusOK,
			schema.Resp{
				Code: schema.NoRoute,
				Msg:  schema.NoRoute.String(),
			},
		)
		return
	}
}
