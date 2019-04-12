package main

import "fmt"

func main()  {
	m := make(map[string]int)

	m["a"] = 23
	fmt.Println(m["a"])

	m["a"] = 32
	fmt.Println(m["a"])

	//delete(m, "a")
	//fmt.Println(m["a"])
	_, isset := m["a"]
	_, isset1 := m["ssss"]

	fmt.Println(isset)
	fmt.Println(isset1)





}
