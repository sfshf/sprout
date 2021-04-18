package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/pkg/jwtauth"
	"github.com/sfshf/sprout/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	SessionIdKey = "sessionId"
)

func SessionIdFromGinX(c *gin.Context) *primitive.ObjectID {
	if sessionId, exists := c.Get(SessionIdKey); exists {
		return sessionId.(*primitive.ObjectID)
	} else {
		return nil
	}
}

func JWT(auth *jwtauth.JWTAuth, repo *repo.Staff) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		if jwtString := c.GetHeader("Authorization"); jwtString != "" {
			subject, err := auth.ParseSubject(jwtString)
			if err != nil {
				JSONWithInvalidToken(c, jwtauth.ErrInvalidToken.Error())
				return
			}
			sessionId, err := primitive.ObjectIDFromHex(subject)
			if err != nil {
				JSONWithInvalidToken(c, err.Error())
				return
			}
			// Verify whether the token is in use, to guarantee an account signed in by only one person.
			err = repo.VerifySignInToken(ctx, &sessionId, &jwtString)
			if err != nil {
				JSONWithInvalidToken(c, err.Error())
				return
			}
			LogWithGinX(c, SessionIdKey, sessionId.Hex())
			c.Set(SessionIdKey, &sessionId)
			c.Next()
			return
		}
	}
}
