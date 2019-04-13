package main

import (
	"math"
	"fmt"
)

type myFloat float64

func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}
//接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法
func main()  {
	f := myFloat(-math.Sqrt(3))
	fmt.Println(f.abs())
	fmt.Println(math.Abs(-0.222))
}