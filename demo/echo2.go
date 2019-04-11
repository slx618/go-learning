package main

import (
	"fmt"
	"os"
)

func main()  {
	//fmt.Println(os.Args[0])
	s := ""
	//fmt.Println(os.Args[0:])
	for arg := 0; arg < len(os.Args); arg ++ {
		s = os.Args[arg]
		fmt.Println(arg, s)
	}

}