package main

import (
	"fmt"
	"sync"
)

//先进先出原则
//引用类型
//初始化后才能使用
//make channel slice map
var ch1 chan int //需要指定通道中的元素类型
var wg sync.WaitGroup

func noBuffChannel() {
	fmt.Println(ch1)
	ch1 = make(chan int) //不带缓冲区的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-ch1
		fmt.Println(x)
	}()
	ch1 <- 10 //没有上面的goroutine会死锁

	ch1 = make(chan int, 10) // 带缓冲区的通道初始化
	fmt.Println(ch1)

	wg.Wait()
}

func buffChannel() {
	fmt.Println(ch1)
	ch1 = make(chan int, 1)
	ch1 <- 10
	// ch <- 10 再存也是阻塞
	//ch1 <- 20 再取就阻塞
	close(ch1)
}

//操作
// 设置值 ch <-
// 取出值 x: = <- ch
// 关闭 close(ch)

func main() {
	buffChannel()
	x := <-ch1
	fmt.Println(x)
}
