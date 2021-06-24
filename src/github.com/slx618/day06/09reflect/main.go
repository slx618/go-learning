package main

import (
	"fmt"
	"reflect"
)

type cat struct {
}

func reflectType(a interface{}) {
	ty := reflect.TypeOf(a) //动态类型
	v := reflect.ValueOf(a) //动态值

	fmt.Println(ty.Kind(), ty.Name(), v.Kind(), v, ty)
}

func reflectSetVal(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
}

func reflectSetVal1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	a := 2.2123
	name := "go 语言"

	reflectType(a)
	reflectType(name)

	var c cat
	reflectType(c)

	b := int64(100)
	reflectSetVal1(&b)
	fmt.Println(b)
}
