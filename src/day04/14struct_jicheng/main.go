package main

import "fmt"

//结构体模拟继承

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Println(a.name, "动了")
}

type dog struct {
	feet uint8
	animal
}

func (d dog) wang() {
	fmt.Println(d.name, "wangwangwang")
}

func main() {
	d := dog{
		feet: 4,
		animal: animal{
			name: "阿黄",
		},
	}

	d.wang()
	d.move()

}
