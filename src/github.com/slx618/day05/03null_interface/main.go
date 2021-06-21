package main

import "fmt"

//interface{} 空接口
//所有的类型都实现了空接口
//任意类型的变量都能保存到空接口中

func show(v interface{}) {
	fmt.Printf("type: %T value: %#v\n", v, v)

}

func main() {

	var m1 map[string]interface{}

	m1 = make(map[string]interface{}, 16)

	m1["name"] = "alex"
	m1["age"] = 18
	m1["s"] = false
	m1["hobby"] = [...]string{"ss", "sss"}

	fmt.Println(m1)

	show(false)
	show(m1)
	show(nil)

}
