package main

import "fmt"

//指针
// & 取地址
// * 根据地址取值
func main() {
	n := 18
	fmt.Println(&n)
	addr := &n
	fmt.Printf("%T\n", addr)
	fmt.Println(addr)
	fmt.Println(*addr)
}
