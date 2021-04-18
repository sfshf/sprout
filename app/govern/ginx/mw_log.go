package ginx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/sfshf/sprout/pkg/json"
	"github.com/sfshf/sprout/pkg/logger"
	"strconv"
	"strings"
	"time"
)

const (
	LoggerEventKey = "logger"
)

func LogWithGinX(c *gin.Context, key, val string) {
	e := LogEventFromGinX(c)
	e = e.Str(key, val)
	c.Set(LoggerEventKey, e)
}

func LogEventFromGinX(c *gin.Context) *zerolog.Event {
	if e, has := c.Get(LoggerEventKey); has {
		return e.(*zerolog.Event)
	}
	return nil
}

func Logger(log *logger.Logger, enable bool) gin.HandlerFunc {
	if !enable {
		return gin.Logger()
	}
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		path := c.FullPath()
		// TODO filtered paths.
		if strings.Contains(path, "/accessLog") {
			c.Next()
			return
		}
		start := time.Now()
		e := log.Info(ctx).
			Str("clientIp", c.ClientIP()).
			Str("proto", c.Request.Proto).
			Str("method", c.Request.Method).
			Str("path", path).
			Str("queries", c.Request.URL.RawQuery).
			Str("requestHeaders", json.Marshal2String(c.Request.Header))
		c.Set(LoggerEventKey, e)
		c.Next()
		var reqBody string
		if byts, has := c.Get(gin.BodyBytesKey); has {
			reqBody = fmt.Sprintf("%s", byts.([]byte))
		}
		var respBody string
		if body, has := c.Get(ResponseBodyKey); has {
			respBody = body.(string)
		}
		e.Str("requestBody", reqBody).
			Str("statusCode", strconv.Itoa(c.Writer.Status())).
			Str("responseHeaders", json.Marshal2String(c.Writer.Header())).
			Str("responseBody", respBody).
			Str("latency", time.Now().Sub(start).String()).
			Msg("")
		return
	}
}
