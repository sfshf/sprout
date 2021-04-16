package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/schema"
)

func IsRoot(rootId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId := SessionIdFromGinX(c)
		if sessionId.Hex() == rootId {
			c.Next()
			return
		} else {
			AbortWithUnauthorized(c, schema.ErrUnauthorized.Error())
			return
		}
	}
}
