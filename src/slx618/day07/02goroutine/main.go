package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)

}

// main 函数启动也会创建一个主 goroutine
//go语言自己实现的线程, go语言的线程切换消耗小 效率高
//goroutine 初始化占用的内存是 2K

//G goroutine 用户态的线程
// M machine 内核线程的虚拟
// P 管理调度 runtime.GOMAXPROCS 最大256 1.5版本后默认是物理线程数
// m:n 把m个goroutine 分配给n个操作系统线程去执行

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
