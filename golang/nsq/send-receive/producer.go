package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
)

type CustomerMessage struct {
	MesID int    `json:"id"`
	Body  string `json:"body"`
}

var nullLogger = log.New(ioutil.Discard, "", log.LstdFlags)
var topic = "msg-test"
var IPStr = "localhost:4150"

func sendMessage() {
	// 1. 创建生产者
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(IPStr, config)
	if err != nil {
		log.Fatalln("链接失败：", err, " ip: ", IPStr)
	}

	// 2. 生成者ping
	errPing := producer.Ping()
	if errPing != nil {
		log.Fatalln("无法ping通: ", errPing)
	}

	// 3. 设置不输出info级别的日志
	producer.SetLogger(nullLogger, nsq.LogLevelInfo)

	// 4. 发布消息
	for i := 0; i < 10; i++ {
		message := "消息发送测试 " + strconv.Itoa(i+10000)
		msg := CustomerMessage{
			MesID: i,
			Body:  message,
		}
		body, _ := json.Marshal(&msg)
		if err := producer.Publish(topic, body); err != nil {
			log.Panic("生产者推送消息失败!")
		}
		log.Println("消息： ", i)
		time.Sleep(time.Millisecond * 500)
	}

	// 5. 生产者停止执行
	producer.Stop()
}

func main() {
	sendMessage()
}
