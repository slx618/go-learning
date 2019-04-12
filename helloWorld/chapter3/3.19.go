package main

import "fmt"

func add() func(int) int  {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

func main()  {
	pos, ng := add(), add()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), // sum 0, 1, 1+2, 3+3, 6+4 ...
			ng(-2*i), // sum 0, -2, -2+-4 -6+-6 ...
		)
	}
}