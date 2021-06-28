package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 10
	close(ch1) //通道关闭可以取值
	//for rst := range ch1 {
	//	fmt.Println(rst)
	//}

	<-ch1
	<-ch1
	x, ok := <-ch1
	fmt.Println(x, ok)

	//可以一直读 值是对应类型的零值
	x, ok = <-ch1
	fmt.Println(x, ok)
}
