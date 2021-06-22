package main

import "fmt"

type cat struct {
}

type dog struct {
}

type person struct {
}

//接口是一种类型
//约束了变量有那些方法
//只要实现了 speak 方法 都是 speaker 类型
//一个变量如果实现了接口中所有的方法 那么这个变量就实现了这个接口
type speaker interface {
	speak() //方法签名
}

func (c cat) speak() {
	fmt.Println("miaomiaomiao")
}

func (d dog) speak() {
	fmt.Println("wangwangwang")
}

func (p person) speak() {
	fmt.Println("aaaaaaaaaaaaa")
}

func da(x speaker) {
	//接收一个参数
	x.speak()
}

func main() {
	c := cat{}
	d := dog{}
	p := person{}

	da(c)
	da(d)
	da(p)

	var ss speaker
	ss = c
	ss.speak()
	ss = d
	ss.speak()
	ss = p
	ss.speak()

}
