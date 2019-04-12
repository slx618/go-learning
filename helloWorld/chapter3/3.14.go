package main

import "fmt"

type myStruct struct {
	l, f float64
}

var s map[string]myStruct

func main()  {
	fmt.Println(s)
	s = make(map[string]myStruct)
	s["ssss"] = myStruct{
		3.23,
		35.2333,
	}
	fmt.Println(s)
	s["SDD"] = myStruct{
		5.66,
		53.222,
	}

	fmt.Println(s)
}
