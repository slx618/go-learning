package main

import (
	"fmt"
	"sort"
)

// 切片的操作是引用传递的 !!!!
// 想要操作独立的切片要用 copy

func main() {

	a1 := []int{1, 2, 3}
	a2 := a1

	//var a3 = make([]int, 0, 3) //长度为0 复制不进去
	var a3 = make([]int, 0, 3)
	copy(a3, a1)

	fmt.Println(a1, a2, a3)
	a1[0] = 999
	fmt.Println(a1, a2, a3)

	//将索引是 2 的删除
	//容量不变
	a1 = append(a1[0:1], a1[2:]...)
	fmt.Println(a1)

	x1 := [...]int{1, 2, 3}
	fmt.Println(x1, len(x1), cap(x1))

	s1 := x1[:]
	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(x1, len(x1), cap(x1))

	var a = make([]int, 5, 10)
	for i := 1; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)

	var b = make([]int, 0, 10)
	for i := 1; i < 10; i++ {
		b = append(b, i)
	}
	fmt.Println(b)

	var c1 = [...]int{3, 5, 7, 7, 89, 0, 976, -1}
	sort.Ints(c1[:]) //对切片进行排序
	fmt.Println(c1)
}
