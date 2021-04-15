package request

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetClientRealIp(c *gin.Context) string {
	if ip := c.GetHeader("X-Real-Ip"); ip != "" {
		return ip
	} else if ip = c.GetHeader("X-Forwarded-For"); ip != "" {
		return ip
	}
	return strings.Split(c.Request.RemoteAddr, ":")[0]
}
