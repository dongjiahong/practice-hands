package models

import (
	"context"

	"sqlent/ent"
)

type UserCountDao struct {
	conn *ent.Client
}

func NewUserCountDao(conn *ent.Client) *UserCountDao {
	return &UserCountDao{conn: conn}
}

func (d *UserCountDao) CreateUserCount(ctx context.Context, u *ent.UserCount) (*ent.UserCount, error) {
	return d.conn.UserCount.
		Create().
		SetSelfBuy(u.SelfBuy).
		SetInviteBuy(u.InviteBuy).
		SetLevel(u.Level).
		Save(ctx)
}
