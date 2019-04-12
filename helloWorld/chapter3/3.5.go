package main

import (
	"fmt"
)

func main()  {
	var a [3]int
	a[0] = 2
	a[2] = 3

	fmt.Println(a)

	myStr := [3]string {"ADSADSA", "21321", "21321SS"}

	var arr1 = a[1:]
	arr2 := myStr[:]

	fmt.Println(arr1)
	fmt.Println(arr2)

	//切片类似引用
	arr1[1] = 11111111111111
	fmt.Println(arr1)
	fmt.Println(a)


}
