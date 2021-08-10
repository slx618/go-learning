package main

import "fmt"

//switch 简化大量if
func main() {

	var n = 5

	switch n {
	case 1:
		fmt.Println(n)
	case 2:
		fmt.Println(n)
	}

	switch z := 12; z {
	case 1, 2, 3, 9:
		fmt.Println(z)
	case 4, 7:
		fmt.Println(n)
	}

	var a = "b"
	switch {
	case a == "a":
		fmt.Println(a)
		fallthrough
	case a == "b":
		fmt.Println(a)
	default:
		fmt.Println("...")
	}
}
