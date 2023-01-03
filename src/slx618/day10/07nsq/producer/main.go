package main

import (
	"bufio"
	"fmt"
	"github.com/nsqio/go-nsq"
	"io"
	"os"
	"strings"
)

//go get github.com/nsqio/go-nsq

var producer *nsq.Producer

func initProducer(str string) (err error) {
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(str, config)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func main() {
	nsqAddr := "127.0.0.1:4102" //tcp 的地址
	err := initProducer(nsqAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println(err)
			continue
		}
		data = strings.TrimSpace(data)

		if strings.ToLower(data) == "q" {
			break
		}

		err = producer.Publish("topic_demo", []byte(data))
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("success:", data)

	}
}
