package mq

import (
	"fmt"

	"github.com/nsqio/go-nsq"
)

type NsqConsumer struct {
	client       *nsq.Consumer
	config       *nsq.Config
	concurrency  int
	topic        string
	channel      string
	nsqlookupdIP string
	nsqdIP       string
	err          error
}

func NewNsqConsumer(topic, channel string) *NsqConsumer {
	return &NsqConsumer{
		config:      nsq.NewConfig(),
		topic:       topic,
		channel:     channel,
		concurrency: 1,
	}
}
func (nc *NsqConsumer) SetMap(options map[string]interface{}) {
	for k, v := range options {
		nc.Set(k, v)
	}
}

func (nc *NsqConsumer) Set(option string, value interface{}) {
	switch option {
	case "topic":
		nc.topic = value.(string)
	case "channel":
		nc.channel = value.(string)
	case "concurrency":
		nc.concurrency = value.(int)
	case "nsqlookupd":
		nc.nsqlookupdIP = value.(string)
	case "nsqd":
		nc.nsqdIP = value.(string)
	default:
		if err := nc.config.Set(option, value); err != nil {
			nc.err = err
		}

	}
}

// Start 消费
func (nc *NsqConsumer) Start(handler nsq.Handler) error {
	if nc.err != nil {
		return nc.err
	}
	client, err := nsq.NewConsumer(nc.topic, nc.channel, nc.config)
	if err != nil {
		return err
	}
	nc.client = client
	nc.client.AddConcurrentHandlers(handler, nc.concurrency)
	return nc.connect()
}

// connect 链接nsq
func (nc *NsqConsumer) connect() error {
	if len(nc.nsqlookupdIP) == 0 && len(nc.nsqdIP) == 0 {
		return fmt.Errorf(`at least one "nsqd" or "nsqlookupd" address must be config`)
	}

	// 有集群就链接集群[优先]
	if len(nc.nsqlookupdIP) > 0 {
		return nc.client.ConnectToNSQLookupd(nc.nsqlookupdIP)
	}

	// 有单节点就连接单节点
	if len(nc.nsqdIP) > 0 {
		return nc.client.ConnectToNSQD(nc.nsqdIP)
	}

	return nil
}

// Stop 停止并等待
func (nc *NsqConsumer) Stop() error {
	nc.client.Stop()
	<-nc.client.StopChan
	return nil
}
