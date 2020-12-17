package main

import (
	"errors"
	"log"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
)

func Sum(args []int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, errors.New("我说他错了")
}

func CallBack(args ...int64) (int64, error) {
	sum := int64(1)
	for _, arg := range args {
		sum *= arg
	}
	return sum, nil
}

func main() {
	cnf, err := config.NewFromYaml("./config.yml", false)
	if err != nil {
		log.Println("config failed ", err)
		return
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		log.Println("start server failed ", err)
		return
	}
	// 注册任务
	if err := server.RegisterTask("sum", Sum); err != nil {
		log.Println("reg task sum failed ", err)
		return
	}
	if err := server.RegisterTask("call", CallBack); err != nil {
		log.Println("reg task  callback failed ", err)
		return
	}

	worker := server.NewWorker("asong", 1)
	go func() {
		if err := worker.Launch(); err != nil {
			log.Println("start worker error ", err)
			return

		}
	}()

	// task signature
	signature1 := &tasks.Signature{
		Name: "sum",
		Args: []tasks.Arg{
			Type:  "[]int64",
			Value: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		RetryTimeout: 100,
		RetryCount:   3,
	}
}
