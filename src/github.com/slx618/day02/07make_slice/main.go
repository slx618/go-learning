package main

import "fmt"

func main() {

	//类型 长度 容量
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v, len=%d cap=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v, len=%d cap=%d\n", s2, len(s2), cap(s1))

	//判断切片为空 len() == 0
	s3 := []int{1, 2, 3}
	s4 := s3
	s3[0] = 9999
	fmt.Println(s3, s4)

	// 切片的遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}

}
