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

// Priorities
// @description Get predefined priority range of roles.
// @id casbin-role-priority
// @tags casbin
// @summary Get predefined priority range of roles.
// @produce json
// @security ApiKeyAuth
// @success 2000 {array} []int "minimum and maximum value of the predefined priority range."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /casbin/priority [GET]
func (a *Casbin) Priorities(c *gin.Context) {
	ginx.JSONWithSuccess(c, []int{model.PriorityMIN, model.PriorityMAX})
	return
}

// AllApiObjActMap
// @description Get an object-action map of all apis.
// @id casbin-role-api-object-action-map
// @tags casbin
// @summary Get an object-action map of all apis.
// @produce json
// @security ApiKeyAuth
// @success 2000 {map} map[string]string "object-action map of all apis."
// @failure 1000 {error} error "feasible and predictable errors."
func (a *Casbin) AllApiObjActMap(c *gin.Context) {
	objActMap := make(map[string]string, len(a.Routes))
	for _, route := range a.Routes {
		objActMap[route.Path] = route.Method
	}
	ginx.JSONWithSuccess(c, objActMap)
	return
}

// AddRole
// @description Add a role.
// @id casbin-role-add
// @tags casbin
// @summary Add a role
func (a *Casbin) AddRole(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.CasbinAddRoleReq
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

func (a *Casbin) AllRoles(c *gin.Context) {
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
