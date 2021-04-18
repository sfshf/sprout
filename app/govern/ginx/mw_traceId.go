package ginx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/pkg/uuid"
)

const (
	TraceIdKey = "traceId"
)

func TraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetHeader(TraceIdKey)
		if traceId == "" {
			uuid, _ := uuid.NewUUID()
			traceId = fmt.Sprintf("%s", uuid)
		}
		LogWithGinX(c, TraceIdKey, traceId)
		c.Writer.Header().Add(TraceIdKey, traceId)
		c.Next()
	}
}
