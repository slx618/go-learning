package main

import (
	"fmt"
	"sync"
)

//1. 启动一个 goroutine 生成100个数 发送到ch1
//2. 启动一个 goroutine 从ch1取值 做平方 放入 ch2
//3/ 在main中打印ch2

var ch1 chan int
var ch2 chan int

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 1; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

//单向通道 只读 <-chan  只写 chan<-
func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok { //什么时候会不 ok? 通道被关闭了 并且值被取完
			break
		}
		ch2 <- x * x
	}
	once.Do(func() {
		close(ch2)
	})
}

var wg sync.WaitGroup
var once sync.Once

func main() {
	ch1 = make(chan int, 100)
	ch2 = make(chan int, 100)
	wg.Add(3)

	go f1(ch1)
	go f2(ch1, ch2)
	go f2(ch1, ch2)
	wg.Wait()

	for rst := range ch2 {
		fmt.Println(rst)
	}

}
