package main

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sfshf/sprout/app/govern/conf"
	"github.com/sfshf/sprout/pkg/jwtauth"
)

func InitAuth() *jwtauth.JWTAuth {
	c := conf.C.JWTAuth
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
