package main

import "fmt"

type myInt int

type person struct {
	string
	int
	myInt
}

func main() {
	p1 := person{"ss", 18, 111}
	fmt.Println(p1.string)
	fmt.Println(p1.int)
	fmt.Println(p1.myInt)
}
