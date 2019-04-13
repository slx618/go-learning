package main

import (
	"math"
	"fmt"
)

type st struct {
	x, y float64
}

func (v st) abs1() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func abs2(v st) float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func main() {
	v1 := st{3, 4}
	v2 := &st{3, 4}

	fmt.Println(v1.abs1())
	fmt.Println(abs2(v1))

	fmt.Println(v2.abs1())
	//fmt.Println(abs2(v2))
	fmt.Println(abs2(*v2))
}
