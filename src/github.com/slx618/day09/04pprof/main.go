package main

import (
	"fmt"
	"time"
)

func f() {
	var c chan int //nil
	for {

		//会循环得很快
		select {
		case c := <-c: //阻塞
			fmt.Println(c)
		default:

		}
	}
}

func main() {
	for i := 0; i < 4; i++ {
		go f()
	}

	time.Sleep(time.Second)
}
