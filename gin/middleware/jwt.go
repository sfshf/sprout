package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/gin/ginx"
	"github.com/sfshf/sprout/pkg/jwtauth"
	"github.com/sfshf/sprout/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	SessionId = "sessionId"
)

func SessionIdFromGinX(c *gin.Context) *primitive.ObjectID {
	if sessionId, exists := c.Get(SessionId); exists {
		return sessionId.(*primitive.ObjectID)
	} else {
		return nil
	}
}

func JWT(auth *jwtauth.JWTAuth) gin.HandlerFunc {
	staffRepo := repo.StaffRepo()
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		if jwtString := c.GetHeader("Authorization"); jwtString != "" {
			subject, err := auth.ParseSubject(jwtString)
			if err != nil {
				ginx.AbortWithInvalidToken(c, jwtauth.ErrInvalidToken.Error())
				return
			}
			sessionId, err := primitive.ObjectIDFromHex(subject)
			if err != nil {
				ginx.AbortWithInvalidToken(c, err.Error())
				return
			}
			// Verify whether the token is in use, to guarantee an account signed in by only one person.
			err = staffRepo.VerifySignInToken(ctx, &sessionId, &jwtString)
			if err != nil {
				ginx.AbortWithInvalidToken(c, err.Error())
				return
			}
			c.Set(SessionId, &sessionId)
			c.Next()
			return
		}
	}
}
