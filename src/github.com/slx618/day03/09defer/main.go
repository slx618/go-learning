package main

import "fmt"

func calc(index string, a, b int) int {
	rst := a + b
	fmt.Println(index, a, b, rst)
	return rst
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) //defer里的func会先执行
	a = 0                                // 上一行 defer 入栈的时候已经用的是a = 1
	defer calc("2", a, calc("20", a, b))
	b = 1
}
