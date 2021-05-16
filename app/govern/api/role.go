package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sfshf/sprout/app/govern/bll"
	"github.com/sfshf/sprout/app/govern/ginx"
	"github.com/sfshf/sprout/app/govern/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Add
// @description Add a new role.
// @id role-add
// @tags role
// @summary Add a new role.
// @accept json
// @produce json
// @param body body bll.AddRoleReq true "required attributes to add a new role."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /role [POST]
func (a *Role) Add(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.AddRoleReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	creator := ginx.SessionIdFromGinX(c)
	if err := a.bll.Add(ctx, creator, &arg); err != nil {
		ginx.JSONWithDuplicateEntity(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

// Authorize
// @description Allocate authorities to a specific role using menu-widgets pairs.
// @id role-allocate-authority
// @tags role
// @summary Allocate authorities to a specific role.
// @accept json
// @produce json
// @param id path string true "id of the role to be allocated authorities."
// @param body body bll.AuthorizeReq true "menu-widgets pairs."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /role/:id/authorize [PUT]
func (a *Role) Authorize(c *gin.Context) {
	ctx := c.Request.Context()
	roleId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	var arg bll.AuthorizeReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.Authorize(ctx, &roleId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

// Evict
// @description Evict a specific role.
// @id role-evict
// @tags role
// @summary Evict a specific role.
// @produce json
// @param id path string true "id of the role to evict."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action"
// @failure 1000 {error} error "feasible and predictable errors."
// @router /role/:id [DELETE]
func (a *Role) Evict(c *gin.Context) {
	ctx := c.Request.Context()
	roleId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.Evict(ctx, &roleId); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

// Update
// @description Update a specific role.
// @id role-update
// @tags role
// @summary Update a specific role.
// @accept json
// @produce json
// @param id path string true "id of the role to evict."
// @param body body bll.UpdateRoleReq true "attributes need to update."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /role/:id [PUT]
func (a *Role) Update(c *gin.Context) {
	ctx := c.Request.Context()
	roleId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	var arg bll.UpdateRoleReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.Update(ctx, &roleId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

// Profile
// @description Get the profile of a role.
// @id role-profile
// @tags role
// @summary Get infos of a role.
// @produce json
// @param id path string true "id of the role."
// @security ApiKeyAuth
// @success 2000 {object} bll.ProfileRoleResp "profile of the role."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /role/:id [GET]
func (a *Role) Profile(c *gin.Context) {
	ctx := c.Request.Context()
	roleId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	// TODO return menu-widgits pairs of a role.
	res, err := a.bll.Profile(ctx, &roleId)
	if err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, res)
	return
}

// List
// @description Get a list of role.
// @id role-list
// @tags role
// @summary Get a list of role.
// @produce json
// @param query query bll.ListRoleReq false "search criteria."
// @security ApiKeyAuth
// @success 2000 {object} bll.ListRoleResp "role list."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /role [GET]
func (a *Role) List(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.ListRoleReq
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

// Enable
// @description Enable or disable a role.
// @id role-enable
// @tags role
// @summary Enable or disable a role.
// @accept json
// @produce json
// @param id path string true "id of the role."
// @param body body bll.EnableRoleReq true "true for enable, or false for disable."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /role/:id/enable [PATCH]
func (a *Role) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.EnableRoleReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	roleId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.Enable(ctx, &roleId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}
