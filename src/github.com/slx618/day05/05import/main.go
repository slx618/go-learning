package main

import (
	"fmt"

	//. fmt 这么引入 fmt.Println 可以简写 Println 但是不建议用

	//别名 & 路径是从 src 开始的
	calcx "github.com/slx618/day05/06calc"
	//不允许导入包但是不使用包, 因为 go 是编译型的 会增加包冗余
	//匿名导入会自动执行 init 只执行 init
	//_ "src"
	//多个包执行顺序 栈机构 先导入 后执行 init
)

func init() {
	fmt.Println("main init")
}

func main() {
	calcx.Add(2, 3)
	fmt.Println("sss")
}
