package main

import "fmt"

type ss struct {
	x, y float64
}

func (v *ss) scale1(f float64)  {
	v.x *= f
	v.y *= f
}

func scale2(v *ss, f float64)  {
	v.x *= f
	v.y *= f
}

func main()  {
	v1 := ss{3, 4}
	v2 := &ss{3, 4}

	v1.scale1(10)
	scale2(&v1, 10)

	v2.scale1(10)
	scale2(v2, 10)

	fmt.Println(v1, v2)
}
