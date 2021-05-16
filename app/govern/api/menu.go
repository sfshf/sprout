package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sfshf/sprout/app/govern/bll"
	"github.com/sfshf/sprout/app/govern/ginx"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Add
// @description Add a new menu, with widgets if has.
// @id menu-add
// @tags menu
// @summary Add a new menu, with widgets if has.
// @accept json
// @produce json
// @param body body bll.AddMenuReq true "required attributes to add a new menu."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /menu [POST]
func (a *Menu) Add(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.AddMenuReq
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
// @description Evict a specific menu.
// @id menu-evict
// @tags menu
// @summary Evict a specific menu.
// @produce json
// @param id path string true "id of the menu to evict."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /menu/:id [DELETE]
func (a *Menu) Evict(c *gin.Context) {
	ctx := c.Request.Context()
	menuId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.Evict(ctx, &menuId); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

// Update
// @description Update a specific menu.
// @id menu-update
// @tags menu
// @summary Update a specific menu.
// @accept json
// @produce json
// @param id path string true "id of the menu to update."
// @param body body bll.UpdateMenuReq true "attributes need to update."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /menu/:id [PUT]
func (a *Menu) Update(c *gin.Context) {

}

// Profile
// @description Get the profile of a menu.
// @id menu-profile
// @tags menu
// @summary Get infos of a menu.
// @produce json
// @param id path string true "id of the menu."
// @security ApiKeyAuth
// @success 2000 {object} bll.ProfileMenuResp "profile of the api."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /menu/:id [GET]
func (a *Menu) Profile(c *gin.Context) {

}

// List
// @description Get a list of menu.
// @id menu-list
// @tags menu
// @summary Get a list of menu.
// @produce json
// @param query query bll.ListMenuReq false "search criteria."
// @security ApiKeyAuth
// @success 2000 {object} bll.ListMenuResp "menu list."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /menu [GET]
func (a *Menu) List(c *gin.Context) {

}

// Enable
// @description Enable or disable a menu.
// @id menu-enable
// @tags menu
// @summary Enable or disable a menu.
// @accept json
// @produce json
// @param id path string true "id of the menu."
// @param body body bll.EnableMenuReq true "true for enable, or false for disable."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /menu/:id/enable [PATCH]
func (a *Menu) Enable(c *gin.Context) {

}
