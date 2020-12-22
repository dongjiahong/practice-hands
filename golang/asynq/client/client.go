package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"

	"asynq-client/tasks"
)

func main() {
	r := asynq.RedisClientOpt{Addr: "localhost:6379"}
	client := asynq.NewClient(r)

	// create a task with typename add payload
	t1 := tasks.NewWelcomeEmailTask(42)
	t2 := tasks.NewReminderEmailTask(42)

	//Process the task immediately.
	res, err := client.Enqueue(t1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result: %+v\n", res)

	// Process the task 24 hours later
	res, err = client.Enqueue(t2, asynq.ProcessIn(24*time.Hour))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result: %+v\n", res)

}
