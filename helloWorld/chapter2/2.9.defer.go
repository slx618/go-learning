package main

import "fmt"

func main()  {

	fmt.Println("counting")

	for i := 0; i < 100; i++ {
		//栈 后进先出
		defer fmt.Println(i)
	}


	fmt.Println("done")
}
