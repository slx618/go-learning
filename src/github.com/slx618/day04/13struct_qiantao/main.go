package main

import "fmt"

//嵌套结构体
type person struct {
	name string
	age  int
	addr address
}

type pp struct {
	name    string
	age     int
	address //匿名嵌套结构体
}

type ppp struct {
	name     string
	age      int
	address  //匿名嵌套结构体
	workAddr //匿名嵌套冲突
}

type workAddr struct {
	province string
	city     string
}

type address struct {
	province string
	city     string
}

func main() {
	p1 := person{
		name: "xxx",
		age:  18,
		addr: address{
			province: "xxx",
			city:     "xxx",
		},
	}

	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.addr.city)
	fmt.Println(p1.addr)
	//fmt.Println(p1.city)

	p2 := pp{
		name: "xxx",
		age:  18,
		address: address{
			"xxx",
			"xxx",
		},
	}
	fmt.Println(p2.city)
	fmt.Println(p2.address.city)

	p3 := ppp{
		name: "xxx",
		age:  18,
		address: address{
			"xxx",
			"xxx",
		},
		workAddr: workAddr{
			"xxx",
			"xxx",
		},
	}
	//fmt.Println(p3.city) //匿名嵌套冲突
	fmt.Println(p3.workAddr.city)

}
