package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "10000"

	// 基于 10 进制
	i, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v  %T\n", i, i)

	// A 字节数组 -> int
	ii, _ := strconv.Atoi(str)
	fmt.Printf("%v  %T\n", ii, ii)

	i1 := 12345
	fmt.Println(string(i1))
	fmt.Println(fmt.Sprintf("%v", i))

	// int -> string
	ss := strconv.Itoa(8)
	fmt.Printf("%v  %T\n", ss, ss)

	//bool
	boolStr := "1"
	b, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%v  %T\n", b, b)

	f := "12.22"
	ff, _ := strconv.ParseFloat(f, 64)
	fmt.Printf("%v  %T\n", ff, ff)

}
