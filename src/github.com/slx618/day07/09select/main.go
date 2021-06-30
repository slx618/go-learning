package main

import (
	"fmt"
	"time"
)

//select 多路复用 随机去做

var ch chan int

func main() {
	ch = make(chan int, 10)
	for {
		rand.Int()
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- rand.Int():
			time.Sleep(time.Second)
		default:
			fmt.Println("hhh")
		}
	}

	//空 select 会一直阻塞
	//select {}
}
