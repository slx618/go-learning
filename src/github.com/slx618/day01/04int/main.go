package main

import "fmt"

func main() {
	var i1 = 10
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) //把十进制转化成二进制
	fmt.Printf("%o\n", i1) //把十进制转化成八进制
	fmt.Printf("%x\n", i1) //把十进制转化成十六进制

	//八进制
	var i2 = 077
	fmt.Printf("%d\n", i2)

	//十六进制
	i3 := 0xf
	fmt.Printf("%d\n", i3)
	//查看变量的类型
	fmt.Printf("%T\n", i3)

	//声明 int8 否则是默认的 int 类型
	i4 := int8(9)
	fmt.Printf("%T\n", i4)

}
