package ginx

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sfshf/sprout/app/govern/schema"
	"github.com/sfshf/sprout/pkg/jwtauth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	SessionIdKey   = "sessionId"
	RedisKeyPrefix = "JWT:"
)

func SessionIdFromGinX(c *gin.Context) *primitive.ObjectID {
	if sessionId, exists := c.Get(SessionIdKey); exists {
		id, _ := primitive.ObjectIDFromHex(sessionId.(string))
		return &id
	} else {
		return nil
	}
}

type TokenVerifier interface {
	TokenExists(ctx context.Context, key string, token string) bool
}

func JWT(auth *jwtauth.JWTAuth, verifier TokenVerifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		if jwtString := c.GetHeader("Authorization"); jwtString != "" {
			subject, err := auth.ParseSubject(jwtString)
			if err != nil {
				JSONWithInvalidToken(c, schema.ErrInvalidToken.Error())
				return
			}
			// Verify whether the token is in use, to guarantee an account signed in by only one person.
			if !verifier.TokenExists(ctx, RedisKeyPrefix+subject, jwtString) {
				JSONWithInvalidToken(c, schema.ErrInvalidToken.Error())
				return
			}
			LogWithGinX(c, SessionIdKey, subject)
			c.Set(SessionIdKey, subject)
			c.Next()
			return
		} else {
			JSONWithInvalidToken(c, schema.ErrInvalidToken.Error())
			return
		}
	}
}
