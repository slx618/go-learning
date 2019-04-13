package main

import (
	"math"
	"fmt"
)

type myStruct struct {
	x, y float64
}

func (v *myStruct) scale(num float64)  {
	v.x = v.x * num
	v.y = v.y * num
}

func (v myStruct) abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func main()  {
	myValue := myStruct{3, 4}
	myValue.scale(100)
	fmt.Println(myValue.abs())

}