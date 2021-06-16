package main

import "fmt"

// 自定义类型
type myInt int

// 类型别名
type youInt = int

func main() {
	var n myInt
	var m youInt
	n = 100
	m = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%T\n", m)

	var c rune // 类型别名
	c = '中'
	fmt.Printf("%T\n", c)

}
