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

	
	BreakTag: // 跳出整个 双层循环
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i == 5 {
				break BreakTag
			} else {
				// fmt.Printf("%d->%d\n", i, j)
			}
		}

	}



	fmt.Println("break done")

	ContinueTag: //可以跳出整个循环 即跳到改标签所定义的行数上
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if j == 5 {
				continue ContinueTag
			} else {
				fmt.Printf("%d->%d\n", i, j)
			}
		}

	}
}
