package main

import "fmt"

// 变量作用域

var x = 100

func f1() {
	//函数中查找变量顺序
	// 1 先在函数中查找
	// 2 找不到往函数外查找 一直找到全局
	var x = 10
	fmt.Println(x)
}

func main() {
	f1()
}
