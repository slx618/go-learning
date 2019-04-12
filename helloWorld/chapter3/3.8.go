package main

import "fmt"

func main()  {
	a := []int{1, 2, 3, 4, 5}

	a = a[:0]
	printSlice(a)

	a = a[0:4]
	printSlice(a)

	a = a[2:]
	printSlice(a)
}

func printSlice(arr []int) {
	fmt.Printf("len:%d cap:%d value:%v\n", len(arr), cap(arr), arr)
}
