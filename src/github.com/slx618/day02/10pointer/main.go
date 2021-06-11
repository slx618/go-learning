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

	//var a *int nil
	//*a = 100

	// new 函数申请一个内存地址 int string 类型的指针
	var a = new(int)
	fmt.Println(a)
	fmt.Println(*a)
	fmt.Println(*a)
	var s = new(string)
	fmt.Println("&s:", s)
	fmt.Println("*s:", *s)

	*a = 100
	fmt.Println(a)
	fmt.Println(*a)

	//make 分配内存 slice  map chan 返回这三种本身

}
