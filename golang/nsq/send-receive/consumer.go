package main

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/nsqio/go-nsq"
)

var topic = "msg-test"
var topicChannel = "ch-msg-test"

var IPStrConsumer = "localhost:4150"

//var IPStrConsumer = "localhost:4061"

type CustomerMessage struct {
	MesID int    `json:"id"`
	Body  string `json:"body"`
}

func deal(msg *nsq.Message) error {
	var cm CustomerMessage
	if err := json.Unmarshal(msg.Body, &cm); err != nil {
		msg.Requeue(time.Second * 2)
		return errors.New("unmarshal err: " + err.Error())
	}
	if cm.MesID%5 == 0 {
		if rand.Intn(2) == 1 {
			msg.Finish()
			return nil
		}
		msg.Requeue(time.Second * 2)
		return errors.New("random err")
	}
	//msg.Finish()
	return nil
}

func doConsumerTask(do nsq.HandlerFunc) {
	// 1. 创建消费者
	config := nsq.NewConfig()
	consumer, errNewCsmr := nsq.NewConsumer(topic, topicChannel, config)
	if errNewCsmr != nil {
		log.Printf("fail to new consumer!, topic=%s, channel=%s", topic, topicChannel)
		return
	}
	// 2. 添加处理消息方法
	consumer.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		log.Printf("message: %v", string(msg.Body))
		return do(msg)
	}))

	// 2. 添加处理消息方法
	//consumer.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
	//log.Printf("message: %v", string(msg.Body))
	//var cm CustomerMessage
	//if err := json.Unmarshal(msg.Body, &cm); err != nil {
	//msg.Requeue(time.Second * 2)
	//return errors.New("unmarshal err: " + err.Error())
	//}
	//if cm.MesID%5 == 0 {
	//if rand.Intn(2) == 1 {
	//msg.Finish()
	//return nil
	//}
	//msg.Requeue(time.Second * 2)
	//return errors.New("random err")
	//}
	////msg.Finish()
	//return nil
	//}))

	// 3. 通过http请求来发现nsqd生产者和配置topic
	if err := consumer.ConnectToNSQD(IPStrConsumer); err != nil {
		log.Panic("ConnectToNSQLookupds can't find nsq")
	}

	// 4. 接收消费者停止通知
	<-consumer.StopChan

	// 5. 获取统计结果
	stats := consumer.Stats()
	log.Printf("message received: %d, finished: %d, requeued: %d, connections: %d\n",
		stats.MessagesReceived, stats.MessagesFinished, stats.MessagesRequeued, stats.Connections)

}
func main() {
	doConsumerTask(deal)
}
