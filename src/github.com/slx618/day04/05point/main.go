package main

import "fmt"

func main() {

	var a int
	a = 100
	b := &a

	fmt.Printf("a %T  b %T \n", a, b)
	fmt.Printf("%p\n", &a) //a 的内存地址
	fmt.Printf("%p\n", b)  // b 的值
	fmt.Printf("%v\n", b)  // b 的值
	fmt.Printf("%p\n", &b) // b 的内存地址
}
