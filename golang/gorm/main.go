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

//type Class struct {
//Id        int
//ClassName string
//Students  []Student
//}

type Id struct {
	Id int
}

type ClassName struct {
	ClassName string
}
type Class struct {
	Id
	ClassName
	Students []Student
}
type C struct {
	Id
	ClassName
	Students []Student
}

func AddStudent(s *Student) *Student {
	fmt.Println(s)
	err := db.Create(s).Scan(&s).Error
	fmt.Println(err)
	return s
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
	//var c []Class
	//err := db.Where("id = 10").Find(&c).Error
	//fmt.Println("1.==> class: ", c)
	//fmt.Println("1.==> err: ", err)

	var cc C
	db = db.Table("class").Where("id = 1").First(&cc) // 指定表名
	fmt.Println("====> name: ", cc.ClassName, " raws: ", db.RowsAffected)

	var ccc Class
	db = db.Table("class").Where("id = 1").First(&ccc) // 指定表名
	fmt.Println("====> name: ", cc.ClassName, " raws: ", db.RowsAffected)

}

func main() {
	initDB()
	//test_raw()

	s := &Student{
		Name: "haha",
	}
	fmt.Println(AddStudent(s))
}
