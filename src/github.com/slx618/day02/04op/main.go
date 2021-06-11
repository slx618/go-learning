package main

import "fmt"

//运算符

func main() {
	var (
		a = 5
		b = 2
	)

	//算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a++ //单独的语句 不能放在 = 号右边
	b--
	fmt.Println(a)
	fmt.Println(b)

	//关系运算符
	fmt.Println(a == b) //相同类型的变量才能比较
	fmt.Println(a != b) //相同类型的变量才能比较
	fmt.Println(a >= b) //相同类型的变量才能比较
	fmt.Println(a <= b) //相同类型的变量才能比较

}
