package main

import (
	"context"
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

func Sum2(args []int64, res int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum + res, nil
}

func SumRetry(args []int64) (int64, error) {
	if rand.Intn(2) == 0 {
		return 0, errors.New("rand err") // 如果只返回错误，那么会按照signature里的配置重试
		//return 0, tasks.NewErrRetryTaskLater("rand err", 2*time.Second) // 出错并直接指定在2秒后执行
	}
	var sum int64 = 0
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}

func Callback(args ...int64) (int64, error) {
	sum := int64(0)
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
	if err := server.RegisterTask("sum2", Sum2); err != nil {
		log.Println("reg task sum failed ", err)
		return
	}
	if err := server.RegisterTask("sumRetry", SumRetry); err != nil {
		log.Println("reg task sumRetry failed ", err)
		return
	}
	if err := server.RegisterTask("callback", Callback); err != nil {
		log.Println("reg task callback failed ", err)
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
		RetryTimeout: 10, // 重试超时
		RetryCount:   3,  // 重试次数
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
		RetryCount:   4, // 重试次数
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

// SimpleGroup 直接执行一组任务
func SimpleGroup() {
	signature1 := &tasks.Signature{
		Name: "sum",
		Args: []tasks.Arg{
			{
				Type:  "[]int64",
				Value: []int64{1, 2, 3},
			},
		},
		RetryCount:   1,
		RetryTimeout: 2,
	}
	signature2 := &tasks.Signature{
		Name: "sum",
		Args: []tasks.Arg{
			{
				Type:  "[]int64",
				Value: []int64{2, 3, 4},
			},
		},
		RetryCount:   1,
		RetryTimeout: 2,
	}
	signature3 := &tasks.Signature{
		Name: "sum",
		Args: []tasks.Arg{
			{
				Type:  "[]int64",
				Value: []int64{3, 4, 5},
			},
		},
		RetryCount:   1,
		RetryTimeout: 2,
	}

	group, err := tasks.NewGroup(signature1, signature2, signature3)
	if err != nil {
		log.Println("add group failed, err: ", err)
		return
	}

	asyncResults, err := server.SendGroupWithContext(context.Background(), group, 3)
	if err != nil {
		log.Println("get async result err: ", err)
		return
	}

	for _, asyncResult := range asyncResults {
		results, err := asyncResult.Get(1)
		if err != nil {
			log.Println(" get result err: ", err)
			continue
		}
		log.Printf("%v %v\n",
			asyncResult.Signature.Args[0].Value,
			tasks.HumanReadableResults(results),
		)
	}
}

// SimpleChord 把一组结果执行完后的结果当作参数发给callback
func SimpleChord() {
	signature := &tasks.Signature{
		Name: "sum",
		Args: []tasks.Arg{
			{
				Type:  "[]int64",
				Value: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		RetryTimeout: 2, // 重试超时
		RetryCount:   4, // 重试次数
	}
	group, err := tasks.NewGroup(signature, signature)
	if err != nil {
		log.Println("new group err: ", err)
		return
	}

	callback := &tasks.Signature{Name: "callback"}
	chord, err := tasks.NewChord(group, callback)
	if err != nil {
		log.Println("new chord err: ", err)
		return
	}

	chordResult, err := server.SendChordWithContext(context.Background(), chord, 0)
	if err != nil {
		log.Println("could not send chord, err: ", err)
		return
	}
	results, err := chordResult.Get(time.Duration(time.Millisecond * 5))
	if err != nil {
		log.Println("get result failed with err: ", err)
		return
	}
	log.Printf("%v", tasks.HumanReadableResults(results))
}

// SimpleChain 链式执行
// 先执行task1然后task2然后task3，当一个任务完成时，结果被附加到chain中
// 下一个任务的参数列表的末尾，最终执行callback
func SimpleChain() {
	signature1 := &tasks.Signature{
		Name: "sum",
		Args: []tasks.Arg{
			{
				Type:  "[]int64",
				Value: []int64{1, 2, 3},
			},
		},
		RetryCount:   1,
		RetryTimeout: 2,
	}
	signature2 := &tasks.Signature{
		Name: "sum2",
		Args: []tasks.Arg{
			{
				Type:  "[]int64",
				Value: []int64{2, 3, 4},
			},
		},
		RetryCount:   1,
		RetryTimeout: 2,
	}
	signature3 := &tasks.Signature{
		Name: "sum2",
		Args: []tasks.Arg{
			{
				Type:  "[]int64",
				Value: []int64{3, 4, 5},
			},
		},
		RetryCount:   1,
		RetryTimeout: 2,
	}

	callback := &tasks.Signature{Name: "callback"}

	chain, err := tasks.NewChain(signature1, signature2, signature3, callback)
	if err != nil {
		log.Println("new chain err: ", err)
		return
	}

	chainResult, err := server.SendChainWithContext(context.Background(), chain)
	if err != nil {
		log.Println("can't send chain, err: ", err)
		return
	}

	result, err := chainResult.Get(time.Duration(time.Millisecond * 50))
	if err != nil {
		log.Println("get chain result err: ", err)
		return
	}
	log.Printf("%v\n", tasks.HumanReadableResults(result))
}
func main() {
	initServer()
	log.Println("simple ========================== go")
	Simple()
	log.Println("simpleretry ========================= go")
	SimpleRetry()
	log.Println("simplegroup ========================= go")
	SimpleGroup()
	log.Println("simplechord ========================= go")
	SimpleChord()
	log.Println("simplechain ========================= go")
	SimpleChain()

}
