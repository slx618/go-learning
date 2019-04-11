package main

import (
	"math"
	"fmt"
)

func pow(x, y, lin float64) float64  {
	if v := math.Pow(x, y); v > lin {
		return v
	}
	//v 在这行不能在使用了
	//类似匿名函数里的变量
	//fmt.Println(v)
	return lin

}

func main()  {
	fmt.Println(pow(2, 3, 10))
	fmt.Println(pow(4, 7, 10))
	fmt.Println(pow(1, 3, 10))
	fmt.Println(pow(7, 3, 10))
}