package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/schema"
)

func MustRootOrSelf(rootId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId := SessionIdFromGinX(c)
		if sessionId.Hex() != rootId && sessionId.Hex() != c.Param("id") {
			JSONWithUnauthorized(c, schema.ErrUnauthorized.Error())
			return
		}
		c.Next()
		return
	}
}
