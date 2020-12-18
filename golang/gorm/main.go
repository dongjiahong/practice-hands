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
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tb_" + defaultTableName
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, dbName)

	db, _ = gorm.Open("mysql", dsn)
	db.SingularTable(true)
}

func test_raw() {
	type Num struct {
		Num float64
	}
	var num Num
	if err := db.Raw("select sum(coin) as num from tb_distribute_record").Scan(&num).Error; err != nil {
		fmt.Println("========> ", err)
	}

	var drs []*DistributeRecord
	if err := db.Raw("select * from tb_distribute_record").Scan(&drs).Error; err != nil {
		fmt.Println("=====>>>> ", err)
	}
	fmt.Println("=====>>>> ", drs)
}

func test_find() {

	var drs []DistributeRecord
	if err := db.Find(&drs).Error; err != nil {
		fmt.Println("====> find err: ", err)
	}
	fmt.Println(drs)

	type DistributeRecordV2 struct {
		DistributeRecord
	}
	var drs2 []DistributeRecordV2
	if err := db.Find(&drs2).Error; err != nil {
		fmt.Println("====> find err: ", err)
	}
	fmt.Println(drs2)
}

func test_struct_raw() {
	var drs []*DistributeRecord
	if err := db.Raw("select id, coin from tb_distribute_record limit 2;").Scan(&drs).Error; err != nil {
		//if err := db.Raw("select id, coin from tb_distribute_record limit 2;").Scan(&drs).Error; err != nil {
		fmt.Println("===> struct err: ", err)
	}
	fmt.Printf("===> drs: %+v", drs[0])
}

func test_struct_raw_parameter(dr interface{}) {
	drs, _ := dr.([]*DistributeRecord)
	if err := db.Raw("select id, coin from tb_distribute_record limit 2;").Scan(&drs).Error; err != nil {
		fmt.Println("===> struct err: ", err)
	}
	fmt.Printf("===> drs: %+v\n", drs[0])
}

func run_test_struct_raw_parameter() {
	drs := []*DistributeRecord{}
	test_struct_raw_parameter(drs)
	fmt.Printf("===> drs: %+v\n", drs)
}

func main() {
	initDB()
	//test_raw()
	//test_find()
	//test_struct_raw()
	run_test_struct_raw_parameter()
}
