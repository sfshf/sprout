package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/schema"
	"net/http"
)

func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(
			http.StatusOK,
			schema.Resp{
				BizCode: schema.NoMethod,
				BizMsg:  schema.NoMethod.String(),
			},
		)
		return
	}
}
