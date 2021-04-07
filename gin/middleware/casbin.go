package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/gin/ginx"
)

func Casbin(enforcer *casbin.SyncedEnforcer, rootSessionId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId := SessionIdFromGinX(c)
		if sessionId == rootSessionId {
			c.Next()
			return
		}
		authorized, err := enforcer.Enforce(sessionId, c.FullPath(), c.Request.Method)
		if err != nil {
			ginx.AbortWithUnauthorized(c, err.Error())
			return
		}
		if !authorized {
			ginx.AbortWithUnauthorized(c, nil)
			return
		}
		c.Next()
		return
	}
}
