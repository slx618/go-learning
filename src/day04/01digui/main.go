package main

import "fmt"

//递归

// 4! = 4*3*2 4*3!
// 5! = 5*4*3*2 5*4!
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

//归纳法
//上台阶
//一次可以走一步 也可以一次走两步 有多少中走法
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return taijie(n-1) + taijie(n-2)
	}
}

var x uint64
var y uint64

func taijie2(n uint64) uint64 {
	for i := uint64(1); i <= n; i++ {
		if n == 1 {
			return 1
		} else if n == 2 {
			return 2
		} else {
			x += taijie2(n - 1)
			y += taijie2(n - 2)
		}
	}

	return x + y
}

func main() {
	fmt.Println(f(30))

	fmt.Println(taijie(10))
	fmt.Println(taijie2(10))
}
