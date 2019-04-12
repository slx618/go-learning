package main

import (
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	myField := strings.Fields(s) //字符串转成数组
	myCount := make(map[string]int)
	for _, v := range myField {
		if _, isset := myCount[v]; isset {
			myCount[v] ++
		} else {
			myCount[v] = 1
		}
	}
	return myCount
}

func main()  {
	s := "I am learning Go!"
	fmt.Println(WordCount(s))
}
