package main

import "fmt"

func f1(f func()) {
	f()
	fmt.Println("f1")
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//定义一个函数对f2 进行包装
func f3(f func(a int, b int), a, b int) func() {
	fmt.Println("f3")

	return func() {
		f(a, b)
	}
}

func main() {
	ff := f2
	f1(f3(ff, 2, 3))

}
