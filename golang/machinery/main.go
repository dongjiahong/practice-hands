package main

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
)

var server *machinery.Server

func Sum(args []int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}

func SumRetry(args []int64) (int64, error) {
	if rand.Intn(1) == 0 {
		return 0, errors.New("rand err")
		//return 0, tasks.NewErrRetryTaskLater("rand err", 2*time.Second) // 出错并在2秒后执行
	}
	var sum int64 = 0
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}

func initServer() {
	// 1.读取文件配置
	cnf, err := config.NewFromYaml("./config.yml", false)
	if err != nil {
		log.Println("config failed ", err)
		return
	}

	// 2.启服务
	server, err = machinery.NewServer(cnf)
	if err != nil {
		log.Println("start server failed ", err)
		return
	}

	// 3.注册任务
	if err := server.RegisterTask("sum", Sum); err != nil {
		log.Println("reg task sum failed ", err)
		return
	}
	if err := server.RegisterTask("sumRetry", SumRetry); err != nil {
		log.Println("reg task sumRetry failed ", err)
		return
	}
	// 4.启动worker
	// 这里的1是限制goruntine的并发数
	worker := server.NewWorker("asong", 1)
	go func() {
		if err := worker.Launch(); err != nil {
			log.Println("start worker error ", err)
			return

		}
	}()
}

// Simple 简单的异步调用，这里加入了延迟执行功能
func Simple() {
	// signature签名，里面包含了调度需要的信息
	// task signature, 通过singature实例传递给server实例来调度任务
	signature := &tasks.Signature{
		Name: "sum",
		Args: []tasks.Arg{
			{
				Type:  "[]int64",
				Value: []int64{1, 2, 3, 4},
			},
		},
		RetryTimeout: 100, // 重试超时
		RetryCount:   3,   // 重试次数
	}
	// 延迟执行，可以不加则不延迟
	eta := time.Now().UTC().Add(time.Second * 2) // 延迟2s执行
	signature.ETA = &eta

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

// SimpleRetry 简单的重试
func SimpleRetry() {
	signature := &tasks.Signature{
		Name: "sumRetry",
		Args: []tasks.Arg{
			{
				Type:  "[]int64",
				Value: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		RetryTimeout: 2, // 重试超时
		RetryCount:   3, // 重试次数
	}

	asyncResult, err := server.SendTask(signature)
	if err != nil {
		log.Fatalln("send task err: ", err)
	}

	res, err := asyncResult.Get(time.Duration(time.Second * 1))
	if err != nil {
		log.Fatalln("get result err: ", err)
	}
	// HumanReadableResults这个方法可以处理反射值，获得最终结果
	log.Printf("get res is %v\n", tasks.HumanReadableResults(res))
}

func main() {
	initServer()
	Simple()
	SimpleRetry()
}
