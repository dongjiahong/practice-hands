package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type DistributeRecord struct {
	Id             int     `json:"id"`
	Coin           float64 `json:"coin"`
	CreateTime     int64   `json:"create_time"`
	Status         int     `json:"status"`
	DistributeCoin float64 `json:"distribute_coin"`
}

func initDB() {
	username := "root"
	password := ""
	host := "localhost"
	port := 3306
	dbName := "db_zkyc_ex"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, dbName)

	db, _ = gorm.Open("mysql", dsn)
}

func main() {
	initDB()

	type Num struct {
		Num float64
	}

	var num Num
	if err := db.Raw("select sum(coin) as num from tb_distribute_record").Scan(&num).Error; err != nil {
		fmt.Println("========> ", err)
	}
	fmt.Println("====> ", num)

	var drs []DistributeRecord
	if err := db.Raw("select * from tb_distribute_record").Scan(&drs).Error; err != nil {
		fmt.Println("=====>>>> ", err)
	}
	fmt.Println("=====>>>> ", drs)
}
