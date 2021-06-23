package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	p := person{
		"jx",
		12,
	}

	str := `{"name":"xxx", "age":12}`

	_ = json.Unmarshal([]byte(str), &p)
	fmt.Println(p)

	s, _ := json.Marshal(p)
	fmt.Println(string(s))

	// orm
	//reflect.TypeOf reflect.ValueOf
}
