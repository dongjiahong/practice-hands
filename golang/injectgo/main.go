package main

import (
	"fmt"
	"github.com/facebookgo/inject"
)

type Conf struct{ Name string } // 根节点
type DB struct{ Name string }   // 根节点

type UserController struct {
	UserService *UserService `inject:""`
	Conf        *Conf        `inject:""`
}

type PostController struct {
	UserService *UserService `inject:""`
	PostService *PostService `inject:""`
	Conf        *Conf        `inject:""`
}

type UserService struct {
	Db   *DB   `inject:""`
	Conf *Conf `inject:""`
}

type PostService struct {
	Db *DB `inject:""`
}

type Server struct { // 顶节点
	UserApi *UserController `inject:""`
	PostApi *PostController `inject:""`
}

func loadConf() *Conf {
	return &Conf{Name: "test-conf"}
}

func connectDB() *DB {
	return &DB{Name: "test-DB"}
}

func main() {
	conf := loadConf() // *Conf
	db := connectDB()  // *DB

	server := Server{}

	graph := inject.Graph{}

	if err := graph.Provide(
		&inject.Object{
			Value: &server,
		},
		&inject.Object{
			Value: conf,
		},
		&inject.Object{
			Value: db,
		},
	); err != nil {
		panic(err)
	}

	if err := graph.Populate(); err != nil {
		panic(err)
	}
	fmt.Println("=> userapi.conf", server.UserApi.Conf)
	fmt.Println("=> ", server.PostApi.UserService.Conf)

}
