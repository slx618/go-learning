package main

import (
	"fmt"
)

var (
	i int
	f float64
	b bool
	s string
)

func main() {

	fmt.Printf("%T,%T,%T,%T\n", i, f, b, s)
	s := "123中文"
	fmt.Printf("有%d个中文\n", check(s))

	chenfa()

}

func chenfa() {

	for y := 1; y <= 9; y++ {
		for x := 1; x <= 9; x++ {
			if x == y {
				fmt.Printf("%vx%v=%v\n", y, x, x*y)
				break
			}
			fmt.Printf("%vx%v=%v ", y, x, x*y)
		}
	}

}

func check(s string) int {
	i := 0
	for _, v := range s {
		if v > 10000 {
			i += 1
		}
	}
	return i
}
