package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/govern/ginx/response"
)

// Casbin return a PERM access control ginx middleware.
func Casbin(enforcer *casbin.SyncedEnforcer, rootSessionId string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId := SessionIdFromGinX(c)
		fmt.Println(sessionId.Hex())
		fmt.Println(rootSessionId)
		if sessionId.Hex() == rootSessionId {
			c.Next()
			return
		}
		// https://casbin.org/docs/en/how-it-works#request
		// A basic request is a tuple object, at least including
		// subject (accessed entity), object (accessed resource) and action (access method).
		// TODO: `sub` should be a role entity.
		sub := sessionId.Hex()
		authorized, err := enforcer.Enforce(sub, c.FullPath(), c.Request.Method)
		if err != nil {
			response.AbortWithUnauthorized(c, err.Error())
			return
		}
		if !authorized {
			response.AbortWithUnauthorized(c, nil)
			return
		}
		c.Next()
		return
	}
}
