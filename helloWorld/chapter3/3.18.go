package main

import (
	"math"
	"fmt"
)

func compute(fn func(float64, float64) float64)	float64 {
	return fn(3, 5)
}

func main()  {
	fc := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(fc(4, 9))
	fmt.Println(compute(fc))
	fmt.Println(compute(math.Pow))
}
