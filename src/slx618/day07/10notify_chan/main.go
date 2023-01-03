package main

import (
	"fmt"
)

func work(id int, jobs <-chan int, rst chan<- int) {
	for v := range jobs {
		rst <- v * v
		notify <- struct{}{}
	}
}

var notify chan struct{} //不占空间

func main() {

	jobs := make(chan int, 100)
	rst := make(chan int, 100)
	notify = make(chan struct{}, 100)

	go func() {
		for i := 1; i <= 99; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	for i := 1; i <= 3; i++ {
		go work(i, jobs, rst)
	}

	go func() {
		for i := 0; i < 99; i++ {
			<-notify
		}
		close(rst)
	}()

	for v := range rst {
		fmt.Println(v)
	}
}
