package main

import "fmt"

func main() {
	// 元素类型为map的切片
	var s1 = make([]map[int]string, 10, 10)

	s1[0] = make(map[int]string, 1)

	s1[0][100] = "A"
	s1[0][101] = "A"
	fmt.Println(s1)

	//值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["key"] = []int{1, 2, 3}
	fmt.Println(m1)
}
