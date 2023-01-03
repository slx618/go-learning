package main

import "fmt"

func main() {
	age := 19
	if age > 18 {
		fmt.Println("老了")
	} else if age > 35 {
		fmt.Println("中年人")
	} else {
		fmt.Println("还年轻")
	}

	if name := "alex"; name != "zz" {
		fmt.Println("name 的作用域只在 if 里")
		fmt.Println(name)
	} else {
		fmt.Println(name)

	}

}
