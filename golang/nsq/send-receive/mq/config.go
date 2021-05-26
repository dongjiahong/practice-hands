package mq

var (
	Topic   = "msg-test"
	Ipstr   = "localhost:4150"
	Channel = "msg-test-ch"
)

type CustomerMessage struct {
	MesID   int    `json:"id"`
	Captcha string `json:"captcha"`
}
