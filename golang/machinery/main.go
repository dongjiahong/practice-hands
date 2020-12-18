package main

import (
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
	return sum, nil
	//return sum, errors.New("我说他错了")
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

	// 启服务
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
	//	if err := server.RegisterTask("call", CallBack); err != nil {
	//		log.Println("reg task  callback failed ", err)
	//		return
	//	}

	// 这里的1是限制goruntine的并发数
	worker := server.NewWorker("asong", 1)
	go func() {
		if err := worker.Launch(); err != nil {
			log.Println("start worker error ", err)
			return

		}
	}()

	// task signature, 通过singature实例传递给server实例来调度任务
	signature := &tasks.Signature{
		Name: "sum",
		Args: []tasks.Arg{
			{
				Type:  "[]int64",
				Value: []int64{1, 2, 3, 4},
			},
		},
		RetryTimeout: 100,
		RetryCount:   1,
	}

	asyncResult, err := server.SendTask(signature)
	if err != nil {
		log.Fatal(" send task err: ", err)
	}

	res, err := asyncResult.Get(time.Duration(time.Second * 1))
	if err != nil {
		log.Fatal("get result err: ", err)
	}
	// HumanReadableResults这个方法可以处理反射值，获得最终结果
	log.Printf("get res is %v\n", tasks.HumanReadableResults(res))
}
