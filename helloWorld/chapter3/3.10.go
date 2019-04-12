package main

import "fmt"

func main()  {
	s := make([]int, 4, 5)
	printSlice(s)

	s1 := make([]int, 5, 54)
	printSlice(s1)
	
}

func printSlice(a []int)  {
	fmt.Println("len=%d, cap=%d, value=$v", len(a), cap(a), a)
}
