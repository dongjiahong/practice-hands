package main

import (
	"context"
	"log"

	"sqlent/ent"
	"sqlent/models"
)

func main() {
	models.Init()
	defer models.Close()

	ctx := context.Background()
	client := models.GetClient()
	userDao := models.NewUserDao(client)

	newU := &ent.User{
		Phone:       "110",
		Password:    "999",
		PID:         1,
		InvitedCode: "8888",
	}
	u, err := userDao.QueryUser(ctx, newU.Phone)
	if err != nil {
		if ent.IsNotFound(err) {
			u, err = userDao.CreateUser(ctx, newU)
			if err != nil {
				panic(err)
			}
			log.Println("create new user: ", u)
		} else {
			panic(err)
		}
	}
	u.Password = "888"
	log.Println("get new user: ", u)
	u, err = userDao.UpdateUser(ctx, u)
	if err != nil {
		panic(err)
	}
	log.Println("update user: ", u)

	userCountDao := models.NewUserCountDao(client)
	uc, err := userCountDao.CreateUserCount(ctx, &ent.UserCount{
		SelfBuy:   12.12,
		InviteBuy: 0,
	})

	u, err = userDao.AddUserCount(ctx, u, uc)
	if err != nil {
		panic(err)
	}
	log.Println("add user count user: ", u)

	us, err := userDao.GetUserWithCount(ctx, u.ID)
	if err != nil {
		panic(err)
	}
	log.Println("get user with count: ", us)
}
