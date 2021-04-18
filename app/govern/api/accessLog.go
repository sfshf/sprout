package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/bll"
	"github.com/sfshf/sprout/app/govern/ginx"
	"github.com/sfshf/sprout/app/govern/schema"
	"go.mongodb.org/mongo-driver/bson"
)

func (a *AccessLog) List(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.AccessLogListReq
	if err := c.ShouldBindQuery(&arg); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	sort := make(bson.M, 0)
	if arg.OrderBy != nil {
		orderBy, err := arg.OrderBy.Values()
		if err != nil {
			ginx.JSONWithInvalidArguments(c, schema.ErrInvalidArguments.Error())
			return
		}
		for k, v := range orderBy {
			sort[k] = v
		}
	}
	res, err := a.bll.List(ctx, &arg, sort)
	if err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, res)
	return
}
