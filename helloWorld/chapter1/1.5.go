package chapter1

import "fmt"

func swap(a, b string) (string, string)  {
	return a, b
}

func main()  {
	a, b := swap("xxxxxx", "ccccccccc")

	fmt.Println(a, b)
}
