package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/ginx"
)

func (a *User) Add(c *gin.Context) {

}

func (a *User) Delete(c *gin.Context) {
	ginx.JSONWithSuccess(c, "DELETE SUCCESS")
}

func (a *User) Update(c *gin.Context) {

}

func (a *User) List(c *gin.Context) {
	ginx.JSONWithSuccess(c, "SUCCESS")
}

func (a *User) Info(c *gin.Context) {

}
