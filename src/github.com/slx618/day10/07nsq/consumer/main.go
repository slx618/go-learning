package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go get github.com/nsqio/go-nsq

type Handle struct {
	Title string
}

var consumer *nsq.Consumer

func (h *Handle) HandleMessage(msg *nsq.Message) error {
	fmt.Printf("%s recv from %v, msg:%v\n", h.Title, msg.NSQDAddress, string(msg.Body))
	return nil
}

func initConsumer(topic, channel, addr string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 3 * time.Second
	consumer, err = nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Println(err)
		return err
	}

	h := &Handle{
		Title: "sss",
	}
	consumer.AddHandler(h)

	// if err := consumer.ConnectToNSQD(addr); err != nil { // 直接连NSQD
	if err := consumer.ConnectToNSQLookupd(addr); err != nil { // 通过lookupd查询
		return err
	}
	return nil
}

func main() {
	nsqAddr := "127.0.0.1:4161" //lookup
	err := initConsumer("topic_demo", "any_channel", nsqAddr)
	//监听一个 channel 生产者那边就会在产生消息的时候往里面放一条数据
	if err != nil {
		fmt.Println(err)
		return
	}
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c
}
