package main

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/Jeffail/tunny"
)

func main() {
	wg := &sync.WaitGroup{}
	pool := tunny.NewFunc(4, func(i interface{}) interface{} {
		defer wg.Done()
		log.Println(i, " goroutine num: ", runtime.NumGoroutine())
		time.Sleep(time.Second)
		return nil
	})
	defer pool.Close()

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go pool.Process(i)
	}
	wg.Wait()
}
