package models

import (
	"context"

	"sqlent/ent"
	"sqlent/ent/user"
)

type UserDao struct {
	conn *ent.Client
}

func NewUserDao(conn *ent.Client) *UserDao {
	return &UserDao{conn: conn}
}

func (d *UserDao) CreateUser(ctx context.Context, u *ent.User) (*ent.User, error) {
	newUser, err := d.conn.User.Create().
		SetPhone(u.Phone).
		SetPassword(u.Password).
		SetPID(u.PID).
		SetInvitedCode(u.InvitedCode).
		Save(ctx)
	return newUser, err
}

func (d *UserDao) QueryUser(ctx context.Context, phone string) (*ent.User, error) {
	return d.conn.User.Query().Where(user.PhoneEQ(phone)).Only(ctx)
}

func (d *UserDao) UpdateUser(ctx context.Context, u *ent.User) (*ent.User, error) {
	return d.conn.User.UpdateOneID(u.ID).
		SetPhone(u.Phone).
		SetPassword(u.Password).
		SetPID(u.PID).
		SetInvitedCode(u.InvitedCode).
		Save(ctx)
}

func (d *UserDao) AddUserCount(ctx context.Context, u *ent.User, uc *ent.UserCount) (*ent.User, error) {
	return d.conn.User.UpdateOneID(u.ID).
		Save(ctx)
}

func (d *UserDao) GetUserWithAll(ctx context.Context, id int64) (*ent.User, error) {
	user, err := d.conn.User.
		Query().
		Where(
			user.IDEQ(id),
		).
		WithCount().
		WithBuyRecord().
		Only(ctx)

	return user, err
}
