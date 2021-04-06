package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/schema"
	"net/http"
)

func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(
			http.StatusOK,
			schema.Resp{
				Code: schema.NoMethod,
				Msg:  schema.NoMethod.String(),
			},
		)
		return
	}
}
