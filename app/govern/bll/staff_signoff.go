package bll

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Staff) SignOff(c *gin.Context, staffId primitive.ObjectID) error {
	return a.staffRepo.DeleteOne(c.Request.Context(), &staffId)
}
