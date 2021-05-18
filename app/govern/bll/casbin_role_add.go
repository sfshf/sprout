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

func (a *Casbin) AddRole(ctx context.Context, req *CasbinAddRoleReq) error {
	_, err := a.enforcer.AddPolicy(
		strings.ToUpper(req.Subject),
		req.Object,
		req.Action,
	)
	return err
}
