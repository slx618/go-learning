package chapter1

import "fmt"

const (
	bigIn uint64 = 1 << 33

	smallInt = bigIn >> 99

	C = 1 << 1 // 2
	D = 2 >> 1
)

func main()  {
	fmt.Println(bigIn, smallInt, C, D)
}
