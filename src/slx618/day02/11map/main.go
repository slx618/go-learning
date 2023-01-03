package main

import "fmt"

func main() {
	//key string 类型 val int 类型
	var m1 map[string]int //会自己申请内存
	fmt.Println(m1 == nil)
	fmt.Printf("len:%d \n", len(m1))

	m1 = make(map[string]int, 10)
	m1["www"] = 123
	m1["age"] = 18

	fmt.Println(m1)
	fmt.Println(m1["xxx"])
	v, ok := m1["xxx"]
	if !ok {
		//如果key不存在 返回的是对应类型的零值
		fmt.Println("无值")
	} else {
		fmt.Println(v)
	}

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	delete(m1, "age")
	delete(m1, "你大爷") // 不存在的key不报错
	for k := range m1 {
		fmt.Println(k)
	}
}
