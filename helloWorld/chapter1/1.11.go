package chapter1

import "fmt"

func main()  {

	var i int = 32
	z := 18
	var s string = "2323232"
	sz := string("sdsadsadsa")
	u := uint(12321321)
	comp := 0.1 + 12i
	var cm2 complex64 = 21 + 2i

	fmt.Printf("type %T, value %v\n", i, i)
	fmt.Printf("type %T, value %v\n", z, z)
	fmt.Printf("type %T, value %v\n", s, s)
	fmt.Printf("type %T, value %v\n", sz, sz)
	fmt.Printf("type %T, value %v\n", u, u)
	fmt.Printf("type %T, value %v\n", comp, comp)
	fmt.Printf("type %T, value %v\n", cm2, cm2)

}
