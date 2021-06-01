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

// https://www.restapitutorial.com/lessons/restquicktips.html

// 200 OK: General success status code. This is the most common code. Used to indicate success.
func JSONWithStatusOK(c *gin.Context, resp *schema.Resp) {
	if resp != nil {
		resp.Msg = http.StatusText(http.StatusOK)
		c.Set(ResponseBodyKey, json.Marshal2String(resp))
	}
	c.JSON(http.StatusOK, resp)
}

// 201 CREATED: Successful creation occurred (via either POST or PUT).
// Set the Location header to contain a link to the newly-created resource (on POST).
// Response body content may or may not be present.
func JSONWithStatusCreated(c *gin.Context, resp *schema.Resp) {
	if resp != nil {
		resp.Msg = http.StatusText(http.StatusCreated)
		c.Set(ResponseBodyKey, json.Marshal2String(resp))
	}
	c.JSON(http.StatusCreated, resp)
}

// 204 NO CONTENT: Indicates success but nothing is in the response body,
// often used for DELETE and PUT operations.
func JSONWithStatusNoContent(c *gin.Context, resp *schema.Resp) {
	if resp != nil {
		resp.Msg = http.StatusText(http.StatusNoContent)
		c.Set(ResponseBodyKey, json.Marshal2String(resp))
	}
	c.JSON(http.StatusNoContent, resp)
}

// 400 BAD REQUEST: General error for when fulfilling the request would cause an invalid state.
// Domain validation errors, missing data, etc. are some examples.
func JSONWithStatusBadRequest(c *gin.Context, resp *schema.Resp) {
	if resp != nil {
		resp.Msg = http.StatusText(http.StatusBadRequest)
		c.Set(ResponseBodyKey, json.Marshal2String(resp))
	}
	c.JSON(http.StatusBadRequest, resp)
	c.Abort()
}

// 401 UNAUTHORIZED: Error code response for missing or invalid authentication token.
func JSONWithStatusUnauthorized(c *gin.Context, resp *schema.Resp) {
	if resp != nil {
		resp.Msg = http.StatusText(http.StatusUnauthorized)
		c.Set(ResponseBodyKey, json.Marshal2String(resp))
	}
	c.JSON(http.StatusUnauthorized, resp)
	c.Abort()
}

// 403 FORBIDDEN: Error code for when the user is not authorized to perform the operation
// or the resource is unavailable for some reason (e.g. time constraints, etc.).
func JSONWithStatusForbidden(c *gin.Context, resp *schema.Resp) {
	if resp != nil {
		resp.Msg = http.StatusText(http.StatusForbidden)
		c.Set(ResponseBodyKey, json.Marshal2String(resp))
	}
	c.JSON(http.StatusForbidden, resp)
	c.Abort()
}

// 404 NOT FOUND: Used when the requested resource is not found, whether it doesn't exist
// or if there was a 401 or 403 that, for security reasons, the service wants to mask.
func JSONWithStatusNotFound(c *gin.Context, resp *schema.Resp) {
	if resp != nil {
		resp.Msg = http.StatusText(http.StatusNotFound)
		c.Set(ResponseBodyKey, json.Marshal2String(resp))
	}
	c.JSON(http.StatusNotFound, resp)
	c.Abort()
}

// 405 METHOD NOT ALLOWED: Used to indicate that the requested URL exists, but the requested HTTP method is not applicable.
// For example, POST /users/12345 where the API doesn't support creation of resources this way (with a provided ID).
// The Allow HTTP header must be set when returning a 405 to indicate the HTTP methods that are supported.
// In the previous case, the header would look like "Allow: GET, PUT, DELETE".
func JSONWithStatusMethodNotAllowed(c *gin.Context, resp *schema.Resp) {
	if resp != nil {
		resp.Msg = http.StatusText(http.StatusMethodNotAllowed)
		c.Set(ResponseBodyKey, json.Marshal2String(resp))
	}
	c.JSON(http.StatusMethodNotAllowed, resp)
	c.Abort()
}

// 409 CONFLICT: Whenever a resource conflict would be caused by fulfilling the request.
// Duplicate entries, such as trying to create two customers with the same information,
// and deleting root objects when cascade-delete is not supported are a couple of examples.
func JSONWithStatusConflict(c *gin.Context, resp *schema.Resp) {
	if resp != nil {
		resp.Msg = http.StatusText(http.StatusConflict)
		c.Set(ResponseBodyKey, json.Marshal2String(resp))
	}
	c.JSON(http.StatusConflict, resp)
	c.Abort()
}

// 500 INTERNAL SERVER ERROR: Never return this intentionally.
// The general catch-all error when the server-side throws an exception.
// Use this only for errors that the consumer cannot address from their end.
func JSONWithStatusInternalServerError(c *gin.Context, resp *schema.Resp) {
	if resp != nil {
		resp.Msg = http.StatusText(http.StatusInternalServerError)
		c.Set(ResponseBodyKey, json.Marshal2String(resp))
	}
	c.JSON(http.StatusInternalServerError, resp)
	c.Abort()
}

func JSONWithInvalidArguments(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		BizCode: schema.InvalidArguments,
		BizMsg:  schema.InvalidArguments.String(),
		Data:    data,
	}
	JSONWithStatusOK(c, resp)
	c.Abort()
}

func JSONWithFailure(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		BizCode: schema.Failure,
		BizMsg:  schema.Failure.String(),
		Data:    data,
	}
	JSONWithStatusOK(c, resp)
	c.Abort()
}

func JSONWithDuplicateEntity(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		BizCode: schema.DuplicateEntity,
		BizMsg:  schema.DuplicateEntity.String(),
		Data:    data,
	}
	JSONWithStatusOK(c, resp)
	c.Abort()
}

func JSONWithInvalidAccountOrPassword(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		BizCode: schema.InvalidAccountOrPassword,
		BizMsg:  schema.InvalidAccountOrPassword.String(),
		Data:    data,
	}
	JSONWithStatusOK(c, resp)
	c.Abort()
}

func JSONWithInvalidToken(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		BizCode: schema.InvalidToken,
		BizMsg:  schema.InvalidToken.String(),
		Data:    data,
	}
	JSONWithStatusOK(c, resp)
	c.Abort()
}

func JSONWithInvalidCaptcha(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		BizCode: schema.InvalidCaptcha,
		BizMsg:  schema.InvalidCaptcha.String(),
		Data:    data,
	}
	JSONWithStatusOK(c, resp)
	c.Abort()
}

func JSONWithUnauthorized(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		BizCode: schema.Unauthorized,
		BizMsg:  schema.Unauthorized.String(),
		Data:    data,
	}
	JSONWithStatusOK(c, resp)
	c.Abort()
}

func JSONWithSuccess(c *gin.Context, data interface{}) {
	resp := &schema.Resp{
		BizCode: schema.Success,
		BizMsg:  schema.Success.String(),
		Data:    data,
	}
	JSONWithStatusOK(c, resp)
}
