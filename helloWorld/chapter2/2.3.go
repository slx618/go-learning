package main

import (
	"fmt"
	"math"
)


func mySqr(x float64) string {
	if x < 0 {
		return mySqr(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main()  {
	fmt.Println(mySqr(-4))
}