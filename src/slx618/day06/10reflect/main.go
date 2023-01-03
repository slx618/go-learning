package main

import (
	"fmt"
	"reflect"
)

type Stu struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	s := Stu{
		"xx",
		18,
	}

	r := reflect.TypeOf(s)
	fmt.Println(r.Name(), r.Kind())

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		fmt.Printf("name: %s, index:%d type:%v jsonTag: %v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	if f, ok := r.FieldByName("Name"); ok {
		fmt.Printf("name: %s, index:%d type:%v jsonTag: %v\n", f.Name, f.Index, f.Type, f.Tag.Get("json"))

	}
}
