package chapter1

import "fmt"

const apple  = "ssss"

func main()  {
	const MyConst  = 1234
	fmt.Printf("type is: %T, value is: %v\n", apple, apple)
	fmt.Printf("type is: %T, value is: %v\n", MyConst, MyConst)
}
