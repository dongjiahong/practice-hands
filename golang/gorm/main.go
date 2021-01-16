package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Student struct {
	Id   int
	Name string
}

type Class struct {
	Id        int
	ClassName string
	Students  []Student
}

func initDB() {
	username := "root"
	password := ""
	host := "localhost"
	port := 3306
	dbName := "db_test"
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
		//return "tb_" + defaultTableName
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, dbName)

	db, _ = gorm.Open("mysql", dsn)
	db.SingularTable(true)
	db.LogMode(true)
}

func test_raw() {
	var c []Class
	err := db.Where("id = 10").Find(&c).Error
	fmt.Println("1.==> class: ", c)
	fmt.Println("1.==> err: ", err)

	var cc Class
	err = db.Where("id = 10").Find(&cc).Error
	fmt.Println("2.==> class: ", cc)
	fmt.Println("2.==> err: ", err)

	var ccc Class
	err = db.Where("id = 10").Scan(&ccc).Error
	fmt.Println("3.==> class: ", ccc)
	fmt.Println("3.==> err: ", err)
}

func main() {
	initDB()
	test_raw()
}
