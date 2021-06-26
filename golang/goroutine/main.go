package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var done = make(chan bool, 1)
var over = make(chan bool, 1)

// 控制并发数量为5，并在完成后通知主程序
func run() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("-----> i: ", i, " time: ", time.Now().Unix(), " go: ", runtime.NumGoroutine())
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
	done <- true
}

func main() {
	run()
	select {
	case <-done:
		if rand.Intn(4) == 2 {
			over <- true
		}
		run()
	case <-over:
		return
	}
	fmt.Println("over")
}
