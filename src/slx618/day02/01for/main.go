package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	for i := 1; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}
