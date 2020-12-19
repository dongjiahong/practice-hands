package main

import (
	"time"

	"github.com/hibiken/asynq"
	"your/app/package/tasks"
)

const redisAddr = "localhost:6379"

func main() {
	r := asynq.RedisClientOpt{Addr: redisAddr}
	c := asynq.NewClient(r)
	defer c.Close()

	// ---------------------------------------------------
	// Example 1: Enqueue task to be processed immediately
	//		Use (*Client).Enqueue method
	// ---------------------------------------------------
	t := tasks.NewEmailDeliveryTask(42, "some:templage:id")
	res, err := 

	// ---------------------------------------------------
	// ---------------------------------------------------
	// ---------------------------------------------------
	// ---------------------------------------------------
	// ---------------------------------------------------
	// ---------------------------------------------------
}
