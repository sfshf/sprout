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
// @router /menus [POST]
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
// @router /menus/:id [DELETE]
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
// @router /menus [GET]
func (a *Menu) List(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.ListMenuReq
	if err := c.ShouldBindQuery(&arg); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	res, err := a.bll.List(ctx, &arg, nil)
	if err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, res)
	return
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
// @router /menus/:id [GET]
func (a *Menu) Profile(c *gin.Context) {
	ctx := c.Request.Context()
	menuId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	res, err := a.bll.Profile(ctx, &menuId)
	if err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, res)
	return
}

// ListWidget
// @description Get a widget list of a specific menu.
// @id menu-profile-widgetList
// @tags menu
// @summary Get a widget list of a specific menu.
// @produce json
// @param id path string true "id of the menu."
// @security ApiKeyAuth
// @success 2000 {object} bll.ListWidgetResp "widget list of a specific menu."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /menus/:id/widgets [GET]
func (a *Menu) ListWidget(c *gin.Context) {
	ctx := c.Request.Context()
	menuId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	res, err := a.bll.ListWidget(ctx, &menuId)
	if err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, res)
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
// @router /menus/:id [PUT]
func (a *Menu) Update(c *gin.Context) {
	ctx := c.Request.Context()
	menuId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	var arg bll.UpdateMenuReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.Update(ctx, &menuId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
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
// @router /menus/:id/enable [PATCH]
func (a *Menu) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	var arg bll.EnableMenuReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	menuId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.Enable(ctx, &menuId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

// AddWidget
// @description Add a widget for a specific menu.
// @id menu-update-widget-add
// @tags menu
// @summary Add a widget for a specific menu.
// @accept json
// @produce json
// @param id path string true "id of the menu."
// @param body body bll.AddWidgetReq true "necessary attributes to add a widget."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /menus/:id/widgets [POST]
func (a *Menu) AddWidget(c *gin.Context) {
	ctx := c.Request.Context()
	menuId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	var arg bll.AddWidgetReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	creator := ginx.SessionIdFromGinX(c)
	if err := a.bll.AddWidget(ctx, creator, &menuId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

// EvictWidget
// @description Evict a widget for a specific menu.
// @id menu-update-widget-evict
// @tags menu
// @summary Evict a widget for a specific menu.
// @accept json
// @produce json
// @param id path string true "id of the menu."
// @param widgetId path string true "id of the widget."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /menus/:id/widgets/:widgetId [DELETE]
func (a *Menu) EvictWidget(c *gin.Context) {
	ctx := c.Request.Context()
	menuId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	widgetId, err := primitive.ObjectIDFromHex(c.Param("widgetId"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.EvictWidget(ctx, &menuId, &widgetId); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}

// ProfileWidget
// @description Get the profile of a widget.
// @id menu-profile-widget-profile
// @tags menu
// @summary Get infos of a widget.
// @produce json
// @param id path string true "id of the menu."
// @param widgetId path string true "id the a widget."
// @security ApiKeyAuth
// @success 2000 {object} bll.ProfileWidgetResp "profile of the widget."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /menus/:id/widgets/:widgetId [GET]
func (a *Menu) ProfileWidget(c *gin.Context) {
	ctx := c.Request.Context()
	menuId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	widgetId, err := primitive.ObjectIDFromHex(c.Param("widgetId"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	res, err := a.bll.ProfileWidget(ctx, &menuId, &widgetId)
	if err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, res)
	return
}

// UpdateWidget
// @description Update infos of a widget.
// @id menu-update-widget-update
// @tags menu
// @summary Update infos of a widget.
// @accept json
// @produce json
// @param id path string true "id of the menu."
// @param widgetId path string true "id the a widget."
// @param body body bll.UpdateWidgetReq true "some attributes to update."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /menus/:id/widgets/:widgetId [PUT]
func (a *Menu) UpdateWidget(c *gin.Context) {
	ctx := c.Request.Context()
	menuId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	widgetId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	var arg bll.UpdateWidgetReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.UpdateWidget(ctx, &menuId, &widgetId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, err.Error())
	return
}

// EnableWidget
// @description Enable or disable a widget.
// @id menu-update-widget-profile
// @tags menu
// @summary Enable or disable a widget.
// @produce json
// @param id path string true "id of the menu."
// @param widgetId path string true "id the a widget."
// @param body body bll.EnableWidgetReq true "true for enable, or false for disable."
// @security ApiKeyAuth
// @success 2000 {null} null "successful action."
// @failure 1000 {error} error "feasible and predictable errors."
// @router /menus/:id/widgets/:widgetId/enable [PATCH]
func (a *Menu) EnableWidget(c *gin.Context) {
	ctx := c.Request.Context()
	menuId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	widgetId, err := primitive.ObjectIDFromHex(c.Param("widgetId"))
	if err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	var arg bll.EnableWidgetReq
	if err := c.ShouldBindBodyWith(&arg, binding.JSON); err != nil {
		ginx.JSONWithInvalidArguments(c, err.Error())
		return
	}
	if err := a.bll.EnableWidget(ctx, &menuId, &widgetId, &arg); err != nil {
		ginx.JSONWithFailure(c, err.Error())
		return
	}
	ginx.JSONWithSuccess(c, nil)
	return
}
