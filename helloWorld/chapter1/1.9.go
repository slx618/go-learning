package chapter1

import (
	"fmt"
	"math/cmplx"
)

var (
	bigInt uint64 = 1 << 60 - 1
	myBool bool = true
	myStr string = "hello world"
	myCom complex128 = cmplx.Sqrt(-5 + 21i)
)

func main()  {
	fmt.Println(bigInt, myBool, myStr, myCom)
}
