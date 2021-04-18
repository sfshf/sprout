package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sfshf/sprout/app/govern/bll"
	"github.com/sfshf/sprout/app/govern/ginx"
	"github.com/sfshf/sprout/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

func (a *Casbin) Priorities(c *gin.Context) {
	ginx.JSONWithSuccess(c, []int{model.PriorityMIN, model.PriorityMAX})
	return
}

func (a *Casbin) ObjectActionMap(c *gin.Context) {
	resource := make(map[string]string, len(a.Routes))
	for _, route := range a.Routes {
		resource[route.Path] = route.Method
	}
	ginx.JSONWithSuccess(c, resource)
	return
}

func (a *Casbin) AddRole(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.AddRoleReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.AddRole(ctx, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Casbin) DeleteRole(c *gin.Context) {
	ctx := c.Request.Context()
	if err := a.bll.DeleteRole(ctx, strings.ToUpper(c.Param("role"))); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Casbin) SetRole(c *gin.Context) {
	ctx := c.Request.Context()
	staffId, err := primitive.ObjectIDFromHex(c.Param("staffId"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.SetRole(ctx, &staffId, c.Param("role")); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Casbin) UnsetRole(c *gin.Context) {
	ctx := c.Request.Context()
	staffId, err := primitive.ObjectIDFromHex(c.Param("staffId"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.UnsetRole(ctx, &staffId, strings.ToUpper(c.Param("role"))); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

func (a *Casbin) Roles(c *gin.Context) {
	ctx := c.Request.Context()
	roles := a.bll.Roles(ctx)
	ginx.JSONWithSuccess(c, roles)
	return
}

func (a *Casbin) StaffsOfRole(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := a.bll.StaffsOfRole(ctx, strings.ToUpper(c.Param("role")))
	if err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, users)
	return
}

func (a *Casbin) RolesOfStaff(c *gin.Context) {
	ctx := c.Request.Context()
	staffId, err := primitive.ObjectIDFromHex(c.Param("staffId"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	roles, err := a.bll.RolesOfStaff(ctx, &staffId)
	if err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, roles)
	return
}

func (a *Casbin) PoliciesOfRole(c *gin.Context) {
	ctx := c.Request.Context()
	policies := a.bll.PoliciesOfRole(ctx, strings.ToUpper(c.Param("role")))
	ginx.JSONWithSuccess(c, policies)
	return
}

func (a *Casbin) UpdatePolicy(c *gin.Context) {

}

func (a *Casbin) RemovePolicy(c *gin.Context) {

}

func (a *Casbin) Policies(c *gin.Context) {

}
