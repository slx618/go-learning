package main

import (
	"math"
	"fmt"
)

type Abser interface {
	Abs() float64
}

type ss1 struct {
	x, y float64
}

type ff float64

func (f ff) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func (v *ss1) Abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func main()  {
	var a Abser
	f := ff(2)
	v := &ss1{3, 4}

	a = f
	fmt.Println(a)
	fmt.Println(a.Abs())

	a = v
	fmt.Println(v)
	fmt.Println(v.Abs())
}
