package main

import (
	"os"
	"fmt"
	"strings"
)

func main()  {
	s := os.Args[1:]
	fmt.Println(strings.Join(s, "参数:"))

}