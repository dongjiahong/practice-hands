package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

type Task struct {
	index int
	nums  []int
	sum   int
	wg    *sync.WaitGroup
}

func (t *Task) Do() {
	for _, num := range t.nums {
		t.sum += num
	}
	t.wg.Done()
}

func taskFunc(data interface{}) {
	task := data.(*Task)
	task.Do()
	fmt.Printf("task: %d sum:%d\n", task.index, task.sum)
}

func main() {
	p, _ := ants.NewPoolWithFunc(10, taskFunc)
	defer p.Release()

	const (
		DataSize    = 10000
		DataPerTask = 100
	)

	nums := make([]int, DataSize)
	rand.Seed(time.Now().Unix())
	for i := range nums {
		x := rand.Intn(1000)
		nums[i] = x
		fmt.Println("-------> ", x, " ------> ", nums[i])
	}

	var wg sync.WaitGroup
	wg.Add(DataSize / DataPerTask)

	tasks := make([]*Task, 0, DataSize/DataPerTask)
	for i := 0; i < DataSize/DataPerTask; i++ {
		task := &Task{
			index: i + 1,
			nums:  nums[i*DataPerTask : (i+1)*DataPerTask],
			wg:    &wg,
		}
		tasks = append(tasks, task)
		p.Invoke(task)
	}

	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())

	var sum int
	for _, task := range tasks {
		sum += task.sum
	}

	var expect int
	for _, num := range nums {
		expect += num
	}

	fmt.Printf("finish all task, result is %d expect: %d\n", sum, expect)
}
