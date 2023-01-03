package main

import "fmt"

//函数类型
func f1() {
	fmt.Println("hello")
}

func f2() int {
	return 10
}

//函数也可以作为参数的类型
func f3(x func()) {
	x()
}

func f4(x, y int) int {
	return x + y
}

//  函数也可以作为返回值
func f5(x func() int) func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret

}

func main() {
	a := f1
	b := f2
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", f4)
	f3(a)
}
