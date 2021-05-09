package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sfshf/sprout/app/govern/bll"
	"github.com/sfshf/sprout/app/govern/ginx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Role) AddRole(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.AddRoleReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	creator := ginx.SessionIdFromGinX(c)
	if err := a.bll.AddRole(ctx, creator, &arg); err != nil {
		ginx.JSONWithDuplicateEntity(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Role) AllocateAuthority(c *gin.Context) {
	ctx := c.Request.Context()
	roleId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	var arg bll.AllocateAuthorityReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.AllocateAuthority(ctx, &roleId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Role) EvictRole(c *gin.Context) {
	ctx := c.Request.Context()
	roleId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.EvictRole(ctx, &roleId); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Role) UpdateRole(c *gin.Context) {
	//ctx := c.Request.Context()
	//var arg bll.UpdateRoleReq
	//if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
	//	ginx.JSONWithInvalidArguments(c, err.Error())
	//	return
	//}

}

func (a *Role) Profile(c *gin.Context) {

}

func (a *Role) List(c *gin.Context) {

}
