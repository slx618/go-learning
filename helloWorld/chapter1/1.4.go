package chapter1

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int  {
	return x + y
}

func main() {
	fmt.Println(add(1, 3))
	fmt.Println(add2(1, 3))
}
