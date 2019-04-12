package main

import (
	"fmt"
)

func main()  {
	var a = [...]string{
		"21",
		"22",
		"22",
		"22",
		"22",
		"22",
	}

	fmt.Println(a)


	var primse = []bool {true, false, true}
	fmt.Println(primse)

	s := []struct {
		a string
		b bool
	}{
		{"xxx", false},
		{"xxx", false},
		{"xxx", false},
		{"xxx", false},
		{"xxx", false},
		{"xxx", false},
	}

	fmt.Println(s)

	s1 := struct{
		x int
	}{
		1,
	}

	fmt.Println(s1)
}
