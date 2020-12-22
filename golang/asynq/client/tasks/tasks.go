package tasks

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	WelcomeEmail  = "email:welcome"
	ReminderEmail = "email:reminder"
)

func NewWelcomeEmailTask(id int) *asynq.Task {
	payload := map[string]interface{}{"user_id": id}
	return asynq.NewTask(WelcomeEmail, payload)
}

func NewReminderEmailTask(id int) *asynq.Task {
	payload := map[string]interface{}{"user_id": id}
	return asynq.NewTask(ReminderEmail, payload)
}

func HandleWelcomeEmailTask(ctx context.Context, t *asynq.Task) error {
	id, err := t.Payload.GetInt("user_id")
	if err != nil {
		return err
	}
	fmt.Printf("Send Welcome Email to User %d\n", id)
	return nil
}

func HandleReminderEmailTask(ctx context.Context, t *asynq.Task) error {
	id, err := t.Payload.GetInt("user_id")
	if err != nil {
		return err
	}
	fmt.Printf("Send Reminder Email to User %d\n", id)
	return nil
}
