package main

import "fmt"

func main() {
	s := [...]int{1, 2, 3, 4, 5, 6, 7}

	s1 := s[1: 2]
	s2 := s[:]
	s3 := s[:3]
	s4 := s[0:]
	//s5 := s[0: 10]

	fmt.Println(s1, s2, s3, s4)

}
