package tasks

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
)

// 任务类型列表
const (
	TypeEmailDelivery = "email:deliver"
	TypeImageResize   = "image:resize"
)

// ------------------------------
// 任务创建函数NewXXXTask
// 一个任务包含一个类型和一个负载
// ------------------------------
func NewEmailDeliveryTask(userID int, tmplID string) *asynq.Task {
	payload := map[string]interface{}{"user_id": userID, "templ_id": tmplID}
	return asynq.NewTask(TypeEmailDelivery, payload)
}

func NewImageResizeTask(src string) *asynq.Task {
	payload := map[string]interface{}{"src": src}
	return asynq.NewTask(TypeImageResize, payload)
}

//------------------------------
// HandleXXXTask函数去处理输入的任务
// 这个函数应当满足asynq.HandlerFunc接口
//
// Handler 不必是一个函数，你也可以定义一个类型
// 只要它满足asynq.Handler接口就行
//------------------------------
func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	userID, err := t.Payload.GetInt("user_id")
	if err != nil {
		return err
	}
	tmplID, err := t.Payload.GetString("template_id")
	if err != nil {
		return err
	}
	fmt.Printf("Send Email to User: user_id = %d, template_id = %s\n", userID, tmplID)
	return nil
}

// ImageProcessor implement asynq.Handler interface.
type ImageProcessor struct {
	// ...fields for struct
}

func (p *ImageProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	src, err := t.Payload.GetString("src")
	if err != nil {
		return err
	}
	fmt.Printf("Resize image: src = %s\n", src)
	// Image resizing code...
	return nil
}

func NewImageProcessor() *ImageProcessor {
	// ...return an instance
	return nil
}
