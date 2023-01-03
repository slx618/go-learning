package main

import "fmt"

type person struct {
	name, gender string
	age          int
}

// 值类型拷贝
func f(x person) { //修改的是副本的 gender
	x.gender = "女"
}

func f2(x *person) { //修改的是副本的 gender
	//(*x).gender = "女"
	x.gender = "女" //语法糖
}

//结构体是值类型
func main() {

	var p person
	p.name = "Alex"
	p.gender = "男"

	f(p)
	fmt.Println(p.gender)

	f2(&p)
	fmt.Println(p.gender)

	//结构体指针 1
	var p2 = new(person)
	p2.name = "lll"
	//(*p2).name = "lll"

	fmt.Println(p2)
	fmt.Printf("%T\n", p2)
	//fmt.Printf("%x\n", p2)
	fmt.Printf("%p\n", p2)  // p2保存的值是内存地址
	fmt.Printf("%p\n", &p2) // p2 的内存地址

	// 结构体指针 2
	var p3 = &person{
		name: "元素",
		age:  18,
	}

	fmt.Printf("%#v", p3)

	// 结构体 key - value 初始化 值要一一对应 并且个数不能少
	p4 := &person{ // person 类型的指针
		"名字",
		"男",
		18,
	}
	fmt.Printf("%#v", p4)

}
