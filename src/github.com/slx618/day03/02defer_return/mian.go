package main

import "fmt"

func main() {

	//go 中的return 不是原子的 而是分为两部分
	// 第一步 返回赋值
	// defer
	// 第二部分 RET指令
	// 如果函数中存在 defer  defer 会在赋值后被调用

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()

	return x
}

func f2() (x int) {

	defer func() {
		x++
	}()

	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()

	return x
}

func f4() (x int) {

	defer func(x int) {
		x++
	}(x)

	return 5
}
