package models

import (
	"context"
	"fmt"

	"sqlent/ent"
	"sqlent/ent/user"
	"sqlent/ent/usercount"
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
		AddCount(uc).
		Save(ctx)
}

func (d *UserDao) GetUserWithCount(ctx context.Context, id int64) ([]*ent.User, error) {
	x, err := d.conn.User.
		Query().
		//WithCount().
		WithCount(func(q *ent.UserCountQuery) {
			q.Select(usercount.FieldLevel)
		}).
		All(ctx)

	for _, u := range x {
		fmt.Printf("===>  u: %+v\n", *u)
		for _, p := range u.Edges.Count {
			fmt.Println("===>> count: ", p)
		}
	}
	return x, err
}
