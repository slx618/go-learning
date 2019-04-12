package main

import "fmt"

func main()  {
	var s []int
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8)
	for i, v := range s {
		fmt.Println(i, v)
	}

	for i := range s {
		s[i] = 1 << uint(i)
	}

	for i := range s {
		s[i] = 1 << uint(i)
	}

	for _, v := range s {
		fmt.Println(v)
	}

}
