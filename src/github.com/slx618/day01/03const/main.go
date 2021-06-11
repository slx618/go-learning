package main

import "fmt"

// 常量
// 定义以后就不能修改了
const pi = 3.141592653

//批量声明常量
const (
	statusOK = 200
	notFound = 404
)

// 如果某一行没有写值 下面几行默认值是第一行的
const (
	n1 = 100
	n2
	n3
)

// iota  const 中iota第一次出现值为 0  每新增一行 iota 增 1
const (
	a1 = iota
	a2 = iota
	a3
)

const (
	b1 = iota // 0
	b2        //iota 1
	_         //iota 2
	b3        // iota 3
)

//插队

const (
	c1 = iota // 0
	c2 = 100
	c3 = iota // 2
	c4        // 3
)

// 多个常量在一行声明
const (
	d1, d2 = iota + 1, iota + 2 // d1 1  d2 2
	d3, d4 = iota + 1, iota + 2 // d3 2 d4 3

)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) //10000000000 = 2^10 二进制到十进制 1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {

	//pi = 23232

	fmt.Println(a1, a2, a3)
	fmt.Println(b1, b2, b3)
	fmt.Println(c1, c2, c3, c4)
	fmt.Println(d1, d2, d3, d4)
}
