package main

import "fmt"

var (
	ui   uint
	ui8  uint8
	u16  uint16
	ui32 uint32
	ui64 uint64

	i   int
	i8  int8
	i16 int16
	i32 int32
	i64 int64

	uptr uintptr // 表示指针
)

func main() {

	var i1 = 10
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) //把十进制转化成二进制
	fmt.Printf("%o\n", i1) //把十进制转化成八进制
	fmt.Printf("%x\n", i1) //把十进制转化成十六进制

	//八进制
	var i2 = 0777
	fmt.Printf("%o\n", i2)
	fmt.Printf("%T\n", i2)

	//十六进制
	i3 := 0xff
	fmt.Printf("%x\n", i3)
	//查看变量的类型
	fmt.Printf("%T\n", i3)

	//声明 int8 否则是默认的 int 类型
	i4 := int8(9)
	fmt.Printf("%T\n", i4)

}
