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
	i := 0
	fmt.Println(i)

	for _, v := range s {
		fmt.Printf("%T\n", v)
		//fmt.Println(n)
		//if n  > 10000 {
		//	i += 1
		//}
	}

	fmt.Println(i)

}
