package bll

import (
	"context"
)

func (a *Casbin) Roles(ctx context.Context) []string {
	return a.enforcer.GetAllSubjects()
}
