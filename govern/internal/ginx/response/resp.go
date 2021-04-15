package response

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/govern/internal/schema"
	"net/http"
)

func AbortWithInvalidArguments(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, schema.Resp{
		Code: schema.InvalidArguments,
		Msg:  schema.InvalidArguments.String(),
		Data: data,
	})
}

func AbortWithFailure(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, schema.Resp{
		Code: schema.Failure,
		Msg:  schema.Failure.String(),
		Data: data,
	})
}

func AbortWithDuplicateEntity(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, schema.Resp{
		Code: schema.DuplicateEntity,
		Msg:  schema.DuplicateEntity.String(),
		Data: data,
	})
}

func AbortWithInvalidAccountOrPassword(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, schema.Resp{
		Code: schema.InvalidAccountOrPassword,
		Msg:  schema.InvalidAccountOrPassword.String(),
		Data: data,
	})
}

func AbortWithInvalidToken(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, schema.Resp{
		Code: schema.InvalidToken,
		Msg:  schema.InvalidToken.String(),
		Data: data,
	})
}

func AbortWithInvalidCaptcha(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, schema.Resp{
		Code: schema.InvalidCaptcha,
		Msg:  schema.InvalidCaptcha.String(),
		Data: data,
	})
}

func AbortWithUnauthorized(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, schema.Resp{
		Code: schema.Unauthorized,
		Msg:  schema.Unauthorized.String(),
		Data: data,
	})
}

func JSONWithSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, schema.Resp{
		Code: schema.Success,
		Msg:  schema.Success.String(),
		Data: data,
	})
}
