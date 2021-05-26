package mq

import (
	"time"

	"github.com/nsqio/go-nsq"
)

type NsqProducer struct {
	procuder *nsq.Producer
}

var defaultNsqProducer *NsqProducer

func InitNsqServer() {
	nsqConfig := nsq.NewConfig()
	producer, err := nsq.NewProducer(Ipstr, nsqConfig)
	if err != nil {
		panic("nsq new producer er: " + err.Error())
	}
	if errPing := producer.Ping(); errPing != nil {
		panic("nsq producer ping err: " + err.Error())
	}
	defaultNsqProducer = &NsqProducer{procuder: producer}
}

func GetNsqProducer() *NsqProducer {
	return defaultNsqProducer
}

func (np *NsqProducer) SendMessage(topic string, message []byte) error {
	return np.procuder.Publish(topic, message)
}

func (np *NsqProducer) SendMessageDefer(topic string, delay time.Duration, message []byte) error {
	return np.procuder.DeferredPublish(topic, delay, message)
}

func (np *NsqProducer) Stop() {
	np.procuder.Stop()
}
