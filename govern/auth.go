package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sfshf/sprout/govern/config"
	"github.com/sfshf/sprout/pkg/jwtauth"
)

func NewAuth() *jwtauth.JWTAuth {
	c := config.C.JWTAuth
	var opts []jwtauth.Option
	if c.Expired > 0 {
		opts = append(opts, jwtauth.SetExpired(c.Expired))
	}
	if c.SigningKey != "" {
		opts = append(opts, jwtauth.SetSigningKey([]byte(c.SigningKey)))
		opts = append(opts, jwtauth.SetKeyFunc(func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwtauth.ErrInvalidToken
			}
			return []byte(c.SigningKey), nil
		}))
	}
	return jwtauth.New(opts...)
}
