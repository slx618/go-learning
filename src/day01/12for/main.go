package main

import "fmt"

func main() {

	//i 作用域 只在 for 中
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

	// 无限循环
	//for {
	//	fmt.Println("infinite")
	//}

	//for range
	s := "wo w我是字符串"
	fmt.Println(len(s))
	for k, v := range s {
		fmt.Printf("%d %c\n", k, v)
	}
}
