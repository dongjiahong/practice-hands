package main

import (
	"log"

	"github.com/hibiken/asynq"

	"asynq-workers/tasks"
)

func main() {
	r := asynq.RedisClientOpt{Addr: "localhost:6379"}
	srv := asynq.NewServer(r, asynq.Config{
		Concurrency: 5,
	})

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.WelcomeEmail, tasks.HandleWelcomeEmailTask)
	mux.HandleFunc(tasks.ReminderEmail, tasks.HandleReminderEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
