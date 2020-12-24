package tasks

import (
	"log"
	"time"
)

// Add 任务
func Add(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}

// LongRunningTask
func LongRunningTask() error {
	log.Println("Long running task started")
	for i := 0; i < 20; i++ {
		log.Println(10 - i)
		time.Sleep(1 * time.Second)
		// 这里可以通过redis更新任务进度
	}

	// 这里可以用redis取存储任务结果， sender可以异步轮询或指定多久后取redis结果
	log.Println("Long running task finished")
	return nil
}
