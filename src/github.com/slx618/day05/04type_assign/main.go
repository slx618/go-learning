package main

import "fmt"

func assign(a interface{}) {
	if str, ok := a.(string); !ok {
		fmt.Println(ok)
	} else {
		fmt.Println(str)
	}

}

func assign2(a interface{}) {

	switch t := a.(type) {
	case string:
		fmt.Println(t)
	case bool:
		fmt.Println(t)
	case int:
		fmt.Println(t)
	}
}

func main() {
	//assign(100)
	assign2("sds")
}
