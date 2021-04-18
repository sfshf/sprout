package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/schema"
	"github.com/sfshf/sprout/pkg/json"
	"net/http"
)

const (
	ResponseBodyKey = "_gin-gonic/gin/response/bodykey"
)

func JSONWithStatusOK(c *gin.Context, resp interface{}) {
	c.Set(ResponseBodyKey, json.Marshal2String(resp))
	c.JSON(http.StatusOK, resp)
}

func JSONWithInvalidArguments(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		Code: schema.InvalidArguments,
		Msg:  schema.InvalidArguments.String(),
		Data: data,
	}
	JSONWithStatusOK(c, resp)
}

func JSONWithFailure(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		Code: schema.Failure,
		Msg:  schema.Failure.String(),
		Data: data,
	}
	JSONWithStatusOK(c, resp)
}

func JSONWithDuplicateEntity(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		Code: schema.DuplicateEntity,
		Msg:  schema.DuplicateEntity.String(),
		Data: data,
	}
	JSONWithStatusOK(c, resp)
}

func JSONWithInvalidAccountOrPassword(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		Code: schema.InvalidAccountOrPassword,
		Msg:  schema.InvalidAccountOrPassword.String(),
		Data: data,
	}
	JSONWithStatusOK(c, resp)
}

func JSONWithInvalidToken(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		Code: schema.InvalidToken,
		Msg:  schema.InvalidToken.String(),
		Data: data,
	}
	JSONWithStatusOK(c, resp)
}

func JSONWithInvalidCaptcha(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		Code: schema.InvalidCaptcha,
		Msg:  schema.InvalidCaptcha.String(),
		Data: data,
	}
	JSONWithStatusOK(c, resp)
}

func JSONWithUnauthorized(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		Code: schema.Unauthorized,
		Msg:  schema.Unauthorized.String(),
		Data: data,
	}
	JSONWithStatusOK(c, resp)
}

func JSONWithSuccess(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		Code: schema.Success,
		Msg:  schema.Success.String(),
		Data: data,
	}
	JSONWithStatusOK(c, resp)
}
