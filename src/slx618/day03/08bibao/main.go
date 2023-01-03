package main

import "fmt"

func calc(base int) (func(int) int, func(int) int) {
	add := func(x int) int {
		base += x
		return base
	}

	sub := func(x int) int {
		base -= x
		return base
	}

	return add, sub
}

func main() {
	f1, f2 := calc(100)
	// base 一直在生效没有初始化成100
	fmt.Println(f1(100), f2(200)) // 100 0
}
