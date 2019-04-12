package main

import "fmt"

func main()  {
	var arr []int

	fmt.Println(len(arr), cap(arr), arr)
	if arr == nil {
		fmt.Println("arr is empty")
	} else {
		fmt.Println(arr)
	}

}
