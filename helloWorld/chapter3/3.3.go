package main

import "fmt"

type myStru1 struct {
	x int
	y int
}

func main()  {
	a := myStru1{2, 4}
	v := &a
	fmt.Println(v.x)
	fmt.Println(v.y)
	fmt.Println((*v).x)

	v.x = 23
	fmt.Println(a.x)
	fmt.Println(v.x)
	a.x = 12
	fmt.Println(a.x)
	fmt.Println(v.x)
}
