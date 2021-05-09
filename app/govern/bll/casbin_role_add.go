package bll

import (
	"context"
	"strings"
)

type CasbinAddRoleReq struct {
	Subject string `json:"subject" binding:"required"`
	Object  string `json:"object" binding:"required"`
	Action  string `json:"action" binding:"oneof=GET HEAD POST PUT PATCH DELETE CONNECT OPTIONS TRACE"`
}

func (a *Casbin) AddRole(ctx context.Context, arg *CasbinAddRoleReq) error {
	_, err := a.enforcer.AddPolicy(
		strings.ToUpper(arg.Subject),
		arg.Object,
		arg.Action,
	)
	return err
}
