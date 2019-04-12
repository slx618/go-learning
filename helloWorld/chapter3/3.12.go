package main

import "fmt"

func main()  {
	var a []int

	printSlice(a)

	a = append(a, 1)
	printSlice(a)


	a = append(a, 2)
	printSlice(a)



	a = append(a, 2, 3, 3, 4, 4, 5, 5)
	printSlice(a) //len=9 cap=10 [1 2 2 3 3 4 4 5 5] ??????????

	s := []int{1, 2, 2, 3, 3, 4, 4, 5, 5}
	printSlice(s)

}


func printSlice(a []int)  {
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
}
