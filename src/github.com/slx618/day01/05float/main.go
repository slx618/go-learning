package main

import (
	"fmt"
)

func main() {
	//math.MaxFloat32 //最大的浮点数

	f1 := 1.23456 //默认是float64
	fmt.Printf("%T\n", f1)

	f2 := float32(234.32)
	fmt.Printf("%T", f2)

	//f2 = f1 //不能赋值
	//f1 = f2

}
