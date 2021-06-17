package main

import "fmt"

//结构体占用一块连续的内存

type x struct {
	a int8 // 8 bit => 1 byte
	b int8
	c int8
	d string
}

func main() {
	m := x{
		a: int8(1),
		b: int8(2),
		c: int8(3),
		d: "xxx",
	}

	fmt.Println(&m.a, &m.b, &m.c, &m.d)
}
