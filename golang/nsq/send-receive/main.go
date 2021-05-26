package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/nsqio/go-nsq"

	"pro-con/mq"
)

func getRandCaptch(limit int) string {
	numeric := [10]byte{'a', 'b', 'c', 'd', 'f', 'e', 'q', 'w', '9', 'x'}
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < limit; i++ {
		fmt.Fprintf(&sb, "%c", numeric[rand.Intn(10)])
	}
	return sb.String()
}

func dealCaptch(captcha string) error {
	fmt.Println("deal captcha: ", captcha)
	if rand.Intn(2) == 0 {
		return fmt.Errorf("rand error")
	}
	return nil
}

func ProducerCaptcha(mesID int) error {
	producer := mq.GetNsqProducer()
	if producer == nil {
		return fmt.Errorf("ProducerCaptcha get producer nil")
	}
	message := mq.CustomerMessage{
		MesID:   mesID,
		Captcha: getRandCaptch(5),
	}
	body, err := json.Marshal(&message)
	if err != nil {
		return err
	}
	return producer.SendMessage(mq.Topic, body)
}

func ConsumerCaptcha(lookupIP, nsqdIP string) {
	consumer := mq.NewNsqConsumer(mq.Topic, mq.Channel)
	consumer.Set("nsqd", nsqdIP)
	consumer.Set("nsqlookupd", lookupIP)
	consumer.Set("max_attempts", 3)
	consumer.Set("max_in_flight", 100)
	consumer.Set("dial_timeout", time.Millisecond*200)

	if err := consumer.Start(nsq.HandlerFunc(func(msg *nsq.Message) error {
		defer func() {
			// Attempts 是重试尝试的次数，HasResponded()表示消息是否被正确消费，只有正确
			// 消费，即返回nil而不是err时，HasResponded() 为true
			fmt.Println("====> attempts: ", msg.Attempts, " response: ", msg.HasResponded())
			if msg.Attempts == 3 {
				fmt.Println("max attempts")
			}
		}()
		var message mq.CustomerMessage
		if err := json.Unmarshal(msg.Body, &message); err != nil {
			fmt.Println("unmarshal msg body err: ", err)
			return err
		}
		if err := dealCaptch(message.Captcha); err != nil {
			fmt.Println("deal err: ", err)
			return err
		}
		msg.Finish()
		return nil
	})); err != nil {
		fmt.Println("ConsumerCaptcha start err: ", err)
	}
}

func main() {
	mq.InitNsqServer()
	ch := make(chan int)
	go ConsumerCaptcha("", mq.Ipstr)
	if err := ProducerCaptcha(1234); err != nil {
		panic(err)
	}
	<-ch
}
