package main

import "fmt"

type myS struct {
	a, b string
	c int
}

type myB struct {
	d string
}

var m = map[string]myS {
	"ss": myS {
		"a", "b", 123,
	},
	"dd": myS{
		"ssss", "sss", 1233,
	},
}

var d = map[string]myB {
	"s": myB {
		"ssss",
	},
}

var c = map[string]myS {
	"a": {"2312321", "2321", 3333},
	"b": {"2312321", "2321", 3333},
	"c": {"2312321", "2321", 3333},
}

func main()  {
	fmt.Println(m)
	fmt.Println(d)
	fmt.Println(c)
	fmt.Println(c["a"])
	fmt.Printf("%T", c["a"])



}
