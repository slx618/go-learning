package chapter1

import "fmt"

func split(sum int) (y, z int)  {
	y = sum + 1
	z = sum + 2
	return
}

func main()  {
	fmt.Println(split(1))
}
