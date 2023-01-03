package main

import "fmt"

func main() {

	deferDemo()
}

func deferDemo() {
	fmt.Println("start")
	//延迟到函数结束前执行
	//先进后出
	defer fmt.Println("hhh")
	defer fmt.Println("hhh1")
	defer fmt.Println("hhh2")
	defer fmt.Println("hhh3")
	fmt.Println("end")

}
