package main

import "fmt"

//切片 slice

func main() {
	var s1 []int //定义一个存放 int 类型的切片 未初始化

	var s2 []string

	fmt.Println(s1, s2)

	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))

	s1 = []int{1, 2, 3}
	s2 = []string{"ss", "3dswwwwwwwwwwwwwwwwwwwwwwwwwww", "ghgf"}

	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	//长度 和容量
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))

	// 由数组得到切片

	a1 := [11]int{1, 9, 2, 32, 32, 32, 3, 21, 2, 13, 1}

	// 左包含右不包含
	s3 := a1[0:4]
	fmt.Println(s3)

	s4 := a1[:4]
	fmt.Println(s4)

	s5 := a1[4:]
	fmt.Println(s5)

	s6 := a1[5:6]
	fmt.Println(s6)

	s7 := a1[:]
	fmt.Println("s7:", s7)
	a1[0] = 99999 //切片是引用类型 修改原数据会引起切片数据变化
	fmt.Println("s7:", s7)

	//切片的容量是底层数组起始位开始计算的容量
	fmt.Println(len(s3), cap(s3))
	fmt.Println(len(s6), cap(s6))

	a1[1] = 999
	//切片再切片 底层数组还是原来那个
	// 修改切片上的数据也会影响原数组
	s8 := s7[1:3]
	s8[0] = 9998
	fmt.Println("s8:", s8)
	fmt.Println("a1:", a1)
	a1[1] = 2099
	fmt.Println("s8:", s8)
	fmt.Println("a1:", a1)

	fmt.Println(len(s8), cap(s8))

}
