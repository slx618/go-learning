package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
	f      func(int) int
	f1     func() int
	f2     func(int, int) int
	f3     func(int, int) map[int]string // 不赋值是零值
}

func main() {
	var Alex person
	Alex.name = "Alex"
	Alex.age = 18
	Alex.gender = "male"
	Alex.hobby = []string{"lol", "dnf"}

	fmt.Println(Alex)
	fmt.Println(Alex.name)
	fmt.Printf("%T\n", Alex)

	//匿名结构体 临时场景
	var s struct {
		name string
		age  int
	}
	s.name = "sss"
	fmt.Printf("%T\n", s)
}
