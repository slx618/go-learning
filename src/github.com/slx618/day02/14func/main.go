package main

import "fmt"

//函数
//把一段逻辑封装起来 直接使用

func sum(x int, y int) (ret int) {
	return x + y
}

// 没有返回值
func sum1(x int, y int) {
	fmt.Println(x + y)
}

func f2() {
	fmt.Println("f2")
}

//没有命名返回值
func f3() int {
	rst := 2
	return rst
}

//命名返回值就相当于已经声明了这个变量
//使用命名返回值 return 可以省略这个变量
func f4() (rst int) {
	rst = 2
	return
}

func f5() (int, string) {
	return 1, "12sv"
}

//多个同类型的参数可以合并类型
func f6(x, y int, m, n string, i, j bool) int {
	return x + y
}

//可变参数
//可变必须放在最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

// go中没有默认参数这个概念
func main() {
	r := sum(1, 2)
	fmt.Println(r)

	_, n := f5()
	fmt.Println(n)

	f7("xxxx", 1, 2, 3, 4, 5, 5)

	x := [3]int{1, 2, 3}
	fmt.Println(x)
	ff(x)
	fmt.Println(x)
}

//
func ff(a [3]int) {
	//go 语言中的函数值传递 都是值传递
	a[1] = 100
}
