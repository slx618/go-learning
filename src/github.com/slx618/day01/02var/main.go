package main

import "fmt"

//var name string
//var age int

//全局 批量声明 声明了可以不使用

var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func main() {
	name = "slx618"
	age = 18
	isOk = true

	//局部变量声明后必须使用
	//var heiheieh string

	fmt.Printf("name: %s, age: %d, isOk: %v", name, age, isOk)

	fmt.Println()

	// 声明变量同时赋值
	var s1 string = "王尼玛"
	fmt.Println(s1)

	//类型推导
	var s2 = "王小丫"
	fmt.Println(s2)

	//简短变量声明 只能在函数内部使用
	s3 := "王大丫"
	fmt.Println(s3)

	// 匿名变量 _

	// 变量不能重复声明

}
