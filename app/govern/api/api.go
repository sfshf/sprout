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
// @description Add a new api.
// @id api-add
// @tags api
// @summary Add a new api.
// @accept json
// @produce json
// @param body body bll.AddApiReq true "required attributes to add a new api."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /api [POST]
func (a *Api) Add(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.AddApiReq
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

// Evict
// @description Evict a specific api.
// @id api-evict
// @tags api
// @summary Evict a specific api.
// @produce json
// @param id path string true "id of the api to evict."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /api/:id [DELETE]
func (a *Api) Evict(c *gin.Context) {
	ctx := c.Request.Context()
	apiId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.Evict(ctx, &apiId); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

// Update
// @description Update a specific api.
// @id api-update
// @tags api
// @summary Update a specific api.
// @accept json
// @produce json
// @param id path string true "id of the api to update."
// @param body body bll.UpdateApiReq true "attributes need to update."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /api/:id [PUT]
func (a *Api) Update(c *gin.Context) {
	ctx := c.Request.Context()
	apiId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	var arg bll.UpdateApiReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.Update(ctx, &apiId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

// Profile
// @description Get the profile of a api.
// @id api-profile
// @tags api
// @summary Get infos of an api.
// @produce json
// @param id path string true "id of the api."
// @security ApiKeyAuth
// @success 2000 {object} bll.ProfileApiResp "profile of the api."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /api/:id [GET]
func (a *Api) Profile(c *gin.Context) {
	ctx := c.Request.Context()
	apiId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	res, err := a.bll.Profile(ctx, &apiId)
	if err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, res)
	return
}

// List
// @description Get a list of api.
// @id api-list
// @tags api
// @summary Get a list of api.
// @produce json
// @param query query bll.ListApiReq false "search criteria."
// @security ApiKeyAuth
// @success 2000 {object} bll.ListApiResp "api list."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /api [GET]
func (a *Api) List(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.ListApiReq
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
// @description Enable or disable an api.
// @id api-enable
// @tags api
// @summary Enable or disable an api.
// @accept json
// @produce json
// @param id path string true "id of the api."
// @param body body bll.EnableApiReq true "true for enable, or false for disable."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /api/:id/enable [PATCH]
func (a *Api) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.EnableApiReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	apiId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.Enable(ctx, &apiId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}
