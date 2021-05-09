package bll

import "context"

type UpdateRoleReq struct {
	Group string `json:"group" binding:""`
	Name  string `json:"name" binding:""`
	Seq   int    `json:"seq" binding:""`
	Icon  string `json:"icon" binding:""`
	Memo  string `json:"memo" binding:""`
}

func (a *Role) UpdateRole(ctx context.Context, arg *UpdateRoleReq) error {
	return nil
}
