package main

import (
	"log"
	"os"
	"time"
)

func main() {
	handle, _ := os.OpenFile("./syslog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	log.SetOutput(handle)
	for {
		log.Println("这是日志....")
		time.Sleep(time.Second * 2)
	}

}
