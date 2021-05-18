package bll

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO return menu-widgits pairs of a role.
type ProfileRoleResp struct {
	Group     string `json:"group,omitempty"`
	Name      string `json:"name,omitempty"`
	Seq       int    `json:"seq,omitempty"`
	Icon      string `json:"icon,omitempty"`
	Memo      string `json:"memo,omitempty"`
	Enable    bool   `json:"enable,omitempty"`
	Creator   string `json:"creator,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	UpdatedAt int64  `json:"updatedAt,omitempty"`
}

func (a *Role) Profile(ctx context.Context, objId *primitive.ObjectID) (*ProfileRoleResp, error) {
	arg, err := a.roleRepo.FindOneByID(ctx, objId)
	if err != nil {
		return nil, err
	}
	res := &ProfileRoleResp{
		Group:     *arg.Group,
		Name:      *arg.Name,
		Seq:       *arg.Seq,
		Icon:      *arg.Icon,
		Memo:      *arg.Memo,
		Enable:    *arg.Enable,
		CreatedAt: int64(*arg.CreatedAt),
	}
	if one, err := a.staffRepo.FindOneByID(ctx, arg.Creator); err != nil {
		return nil, err
	} else {
		res.Creator = *one.Account
	}
	if arg.UpdatedAt != nil {
		res.UpdatedAt = int64(*arg.UpdatedAt)
	}
	return res, nil
}
