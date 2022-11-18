package main

import (
	"fmt"
)

//切片 slice 引用类型
//扩容策略
//如果申请的容量大于原来的两倍 直接扩容至新申请的容量
//如果小于1024 直接两倍
//如果大于1024 每次增加四分之一
//具体存储的值类型不同 扩容策略也有一定的不同
// 切片的操作是引用传递的 !!!!
func main() {
	var s1 []int //定义一个存放 int 类型的切片 没有分配内存 未初始化

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
	fmt.Printf("s1 len(%d) cap(%d)\n", len(s1), cap(s1))
	fmt.Printf("s2 len(%d) cap(%d)\n", len(s2), cap(s2))

	// 由数组得到切片

	a1 := [11]int{1, 9, 2, 32, 32, 32, 3, 21, 2, 13, 1}
	fmt.Printf("a1:%T\n", a1)

	//左包含右不包含
	// 左闭右开
	s3 := a1[0:4]
	fmt.Println("s3:", s3)

	s4 := a1[:4]
	fmt.Println("s4:", s4)

	s5 := a1[4:]
	fmt.Println("s5:", s5)

	s6 := a1[5:6]
	fmt.Printf("s6:%v %T\n", s6, s6)

	s7 := a1[:]
	fmt.Println("s7:", s7)
	a1[0] = 99999 //切片是引用类型 修改原数据会引起切片数据变化
	fmt.Println("s7:", s7)

	s9 := s3[1:]
	fmt.Println("s9:", s9)

	//切片的容量是从底层数组起始位开始计算容量 还要计算切片最右边没有被包含的那部分容量
	fmt.Printf("s3 len(%d) cap(%d)\n", len(s3), cap(s3))
	fmt.Printf("s6 len(%d) cap(%d)\n", len(s6), cap(s6))
	fmt.Printf("s9 len(%d) cap(%d)\n", len(s9), cap(s9))

	a1[1] = 999
	//切片再切片 底层数组还是原来那个
	//修改切片上的数据也会影响原数组
	s8 := s7[1:3]
	s8[0] = 9998
	fmt.Println("s8:", s8)
	fmt.Println("a1:", a1)
	a1[1] = 2099
	fmt.Println("s8:", s8)
	fmt.Println("a1:", a1)

	fmt.Println(len(s8), cap(s8))

}
