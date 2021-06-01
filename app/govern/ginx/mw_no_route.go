package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/schema"
	"net/http"
)

func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(
			http.StatusOK,
			schema.Resp{
				BizCode: schema.NoRoute,
				BizMsg:  schema.NoRoute.String(),
			},
		)
		return
	}
}
