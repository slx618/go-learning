package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var z = 10
	for ; z > 0; z-- {
		fmt.Println(z)
	}

	var x = 5
	for x > 0 {
		fmt.Println(x)
		x--
	}

	//for {
	//	fmt.Println("infinite")
	//}

	//for range
	s := "wo w我是字符串"

	for k, v := range s {
		fmt.Printf("%d %c\n", k, v)
	}
}
