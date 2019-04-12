package main

import (
	"math"
	"fmt"
)

type myS struct {
	x, y float64
}

func abs(v myS) float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func main()  {
	a := myS{3, 4}
	fmt.Println(abs(a))
}


