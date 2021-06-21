package main

import "fmt"

//二进制实例用途
const (
	eat   int = 4
	sleep int = 2
	play  int = 1
)

// 111
// 吃饭 100
// 睡觉 010
// 打游戏 001
func f(arg int) {
	fmt.Printf("%b\n", arg)
}

func main() {
	f(eat | play)
	f(eat | play | sleep)
}
