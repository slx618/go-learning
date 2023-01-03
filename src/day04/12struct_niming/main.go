package main

import "fmt"

//自定义类型
type myInt int //打印是 main.myInt

//类型别名
type youInt = int //打印是 int

//字段名匿名
type person struct {
	string
	int
	myInt
	youInt
}

func main() {
	p1 := person{"ss", 18, 111, 222}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
	fmt.Println(p1.myInt)
	fmt.Println(p1.youInt)

	p2 := struct {
		int
		string
	}{1, "23"}

	p3 := struct {
		age int
	}{18}

	fmt.Println(p2, p3)
}
