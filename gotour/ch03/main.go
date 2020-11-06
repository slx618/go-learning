package main

import "fmt"

func main() {
	i := 10
	
	if i > 10 {
		fmt.Println("i>10")

	} else {
		fmt.Println("i<=10")
	}

	i = 6
	if i > 10 {
		fmt.Println("i>10")
	} else if i > 5 && i <= 10 {
		fmt.Println("5<i<=10")
	} else {
		fmt.Println("i<=5")
	}

	switch i := 6; {
		case i > 10:
		case i > 5 && i <= 10:
		default:
	}
}
