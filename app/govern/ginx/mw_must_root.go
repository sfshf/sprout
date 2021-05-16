package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/schema"
)

func MustRoot(rootId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId := SessionIdFromGinX(c)
		if sessionId.Hex() != rootId {
			JSONWithUnauthorized(c, schema.ErrUnauthorized.Error())
			return
		}
		c.Next()
		return
	}
}
