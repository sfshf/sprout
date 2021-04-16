package bll

import (
	"context"
)

func (a *Casbin) PoliciesOfRole(ctx context.Context, role string) [][]string {
	return a.enforcer.GetFilteredPolicy(0, role)
}
