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
	// 1.查询用户
	u, err := userDao.QueryUser(ctx, newU.Phone)
	if err != nil {
		if ent.IsNotFound(err) {
			// 1.1创建用户
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
	// 2.更新用户密码
	u, err = userDao.UpdateUser(ctx, u)
	if err != nil {
		panic(err)
	}
	log.Println("update user: ", u)

	userCountDao := models.NewUserCountDao(client)
	// 3.创建用户等级信息
	uc, err := userCountDao.CreateUserCount(ctx, u, &ent.UserCount{
		SelfBuy:   12.12,
		InviteBuy: 0,
	})

	// 3.1.给用户添加等级
	u, err = userDao.AddUserCount(ctx, u, uc)
	if err != nil {
		panic(err)
	}
	log.Println("add user count user: ", u)

	// 4.添加用户购买记录
	userBuyRecordDao := models.NewUserBuyRecordDao(client)
	ubr, err := userBuyRecordDao.CreateUserBuyRecord(ctx, u, &ent.UserBuyRecord{
		Power:      1,
		PowerNum:   3,
		TotalPower: 3.0,
		TotalDay:   120,
		RemainDay:  120,
		Node:       "f10101",
		UsedUsdt:   123.444111,
		BuyDate:    "2022-01-02 12:21:11",
	})
	if err != nil {
		panic(err)
	}
	log.Println("create user buy record ubr: ", ubr)

	// 5.查询用户
	us, err := userDao.GetUserWithAll(ctx, u.ID)
	if err != nil {
		panic(err)
	}
	log.Println("get user all info: ", us.Edges)
}
