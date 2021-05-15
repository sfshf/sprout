package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/bll"
	"github.com/sfshf/sprout/app/govern/ginx"
	"github.com/sfshf/sprout/app/govern/schema"
)

func (a *AccessLog) List(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.AccessLogListReq
	if err := c.ShouldBindQuery(&arg); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	sort, err := schema.OrderByToBsonM(arg.OrderBy)
	if err != nil {
		ginx.JSONWithInvalidArguments(c, schema.ErrInvalidArguments.Error())
		return
	}
	res, err := a.bll.List(ctx, &arg, sort)
	if err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, res)
	return
}
