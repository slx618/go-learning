package main

import "fmt"

type myStru struct {
	x, y int
}

func main()  {
	x := myStru{2, 3}
	fmt.Println(myStru{})
	x.x = 23232
	x.y = 332122
	fmt.Println(x.x)
	fmt.Println(x.y)

}
