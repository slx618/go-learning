package main

import "fmt"

//不能给别的包的类型添加方法 只能给自己包里类型添加方法
type myInt int

func (m myInt) hello() {
	fmt.Println("hello")
}

func main() {
	m := myInt(10)
	m.hello()
}
