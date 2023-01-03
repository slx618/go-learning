package main

import (
	"fmt"
	"os"
)

//1 文件对象类型
//2 获取文件对象详细信息
func main() {
	handle, _ := os.Open("./main.go")
	fmt.Printf("%T\n", handle)

	info, _ := handle.Stat()
	fmt.Println(info.Size())
}
