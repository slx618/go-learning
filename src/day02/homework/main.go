package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "ni da da da ye ye e e ni de ee e e !"
	countWords(s)
}

func countWords(s string) {
	strSlice := strings.Split(s, " ")
	fmt.Println(strSlice)

	var m1 = make(map[string]int, len(strSlice))
	for _, v := range strSlice {
		if _, ok := m1[v]; ok {
			m1[v]++
		} else {
			m1[v] = 1
		}

	}
	fmt.Println(m1)
}
