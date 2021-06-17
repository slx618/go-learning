package main

import "fmt"

//问题 1 myInt(100) ???
type myInt int

// 问题 2 结构体初始化

type person struct {
	name string
	age  int
}

func main() {
	// 声明一个 int32 的 10
	var x = int32(10)
	//var x int32
	//x = 10
	//var x int32 = 10

	fmt.Println(x)

	var m myInt
	m = 100

	m1 := myInt(100)
	fmt.Println(m)
	fmt.Println(m1)

	var p person
	p.name = "sss"
	p.age = 18
	fmt.Println(p)

	var p1 person
	p1.name = "sss"
	p1.age = 18
	fmt.Println(p1)

	var p2 = person{
		name: "xxx1",
		age:  18,
	}
	fmt.Println(p2)

	p3 := person{
		"xxx",
		18,
	}

	fmt.Println(p3)

}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}
