package main

import "fmt"

func fibonacci() func(int) int {
	f1, f2, f3 := 0, 0, 0

	return func(x int) int {
		if x == 1 || x == 2 {
			f1, f2 = 1, 1
			return 1
		}
		f3 = f2 + f1 // 2:1 + 1 = 2 3: f2 = 2 f1=1
		f1 = f2
		f2 = f3
		return f3
	}
}

func main()  {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
