package main

import "fmt"

func main()  {

	var i, j = 1, 2


	p := &i
	//fmt.Println(p)
	//fmt.Println(*p)
	//fmt.Println(i)
	*p = 123222222222222321
	//fmt.Println(p)
	//fmt.Println(*p)
	//fmt.Println(i)

	*p = j
	fmt.Println(j)
	*p = *p / 2
	fmt.Println(i)
	fmt.Println(j)

}
