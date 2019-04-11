package main

import "fmt"

type myStru4 struct{
	x int
	y int
}

var (
	z = myStru4{}
	p = &z
	s = &myStru4{2, 4}
	v = myStru4{1, 3}


)

func main()  {
	fmt.Println(z, p, s, v)


}
