package main

import "fmt"

func main() {

	for i := 'Z'; i > 'A'; i-- {
		for j := 1; j < 10; j++ {
			if i == 'C' {
				goto xx
			}
			//fmt.Printf("%c->%d\n", i, j)

		}

	}

	// label 标签
xx:
	fmt.Println("done")

BreakTag:
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i == 5 {
				break BreakTag
			} else {
				//fmt.Printf("%d->%d\n", i, j)
			}
		}

	}

ContinueTag: //跳过外层的循环用的
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i == 5 {
				continue ContinueTag
			} else {
				fmt.Printf("%d->%d\n", i, j)
			}
		}

	}
}
