package main

import (
	"math"
	"fmt"
)

type s1 struct {
	x, y float64
}

//指针接收值 能够直接修改器接收着指向的值
func (v *s1) scale(f float64)  {
	v.x *= f
	v.y *= f
}



func main()  {
	v := s1{3, 4}
	v.scale(10)
	fmt.Println(v)
	fmt.Printf("before scale %+v abs %v\n", v, v.abs())
	v.scale(10)
	fmt.Printf("after scale %+v abs %v\n", v, v.abs())

}

//避免在每次调用方法是复制该值, 若值的类型为大型结构体时 这样做更高效
func (v *s1) abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}
