package main

import (
	"math"
	"fmt"
)

type s struct {
	x, y float64
}

func scale(v *s, f float64)  {
	v.x = v.x * f
	v.y = v.y * f
}

func abs(v s) float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func main()  {
	myV := s{3, 4}
	asd := &s{32, 69}
	scale(&myV, 1000)
	scale(asd, 1000)

	fmt.Println(abs(myV))

	fmt.Println(asd)
}