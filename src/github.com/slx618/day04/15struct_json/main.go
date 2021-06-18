package main

import (
	"encoding/json"
	"fmt"
)

//结构体和 json

//结构体和 json 互相转换

//字段要大写 否则 json 包访问不到
type person struct {
	Name string `json:"name" db:"name" ini:"name"` //tag
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "xxx",
		Age:  18,
	}

	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", string(b))
		fmt.Println(string(b))
	}
	v := string(b)
	var p2 person
	json.Unmarshal([]byte(v), &p2) // &p2 是为了能在这个函数里能给修改这个 p2
	fmt.Printf("%#v\n", p2)
	fmt.Println(p2.Age)

}
