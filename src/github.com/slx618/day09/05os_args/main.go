package main

import (
	"fmt"
	"os"
)

//获取命令行函数
func main() {
	fmt.Println(os.Args)
	fmt.Printf("%T\n", os.Args)
}
