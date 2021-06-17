package main

import "fmt"

type person struct {
	name string
	age  int
}

type dog struct {
	name string
}

//构造函数 约定成俗 以 new 开头
//返回结构体 还是结构体指针
//当结构体比较大的时候选择返回指针 减少内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func main() {
	p1 := newPerson("alex", 18)
	p2 := newPerson("xxx", 19)

	d1 := newDog("xxx")
	fmt.Println(p1, p2, d1)
}
