package main

import "fmt"

//底层原理
//1. 函数可以作为返回值
//2. 函数内部查找变量的顺序 先在自己内部找 找不到王外层找
func f1(f func()) {
	f()
	fmt.Println("f1")
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//定义一个函数对f2 进行包装
func f3(f func(int, int), a, b int) func() {
	fmt.Println("f3")

	return func() {
		f(a, b)
	}
}

func addr() func(int) int {
	x := 100
	return func(y int) int {
		return x + y
	}
}

func main() {
	ff := f2
	f1(f3(ff, 2, 3))

	f1 := addr()
	fmt.Println(f1(200))

}
