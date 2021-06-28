package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)

}

// main 函数启动也会创建一个主 goroutine
func main() {
	//开启一个单独的 goroutine
	for i := 1; i < 100; i++ {
		//go hello(i)
		go func() {
			//取外部变量 打印的时候 i 已经改变了
			fmt.Println("niming", i)
		}()
	}
	// 函数结束了 main 函数启动的 goroutine 也结束了
	fmt.Println("main")

	time.Sleep(time.Second)

}
