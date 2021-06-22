package main

//package 关键字 包
//一个文件夹就是一个包
//标识符指的是 变量名 函数名 结构体 接口名 常量
//包的外面只能读取到首字母大写命名的func
//包导入的路径是从 $GOPAHT/src 开始找的
import (
	"fmt"

	//. fmt 这么引入 fmt.Println 可以简写 Println 但是不建议用

	//别名 & 路径是从 src 开始的
	calc "github.com/slx618/day05/06calc"
	//不允许导入包但是不使用包, 因为 go 是编译型的 会增加包冗余
	//匿名导入会自动执行 init 只执行 init
	//_ "src"
	//多个包执行顺序 栈机制 先导入 后执行 init
)

//一个包里只有一个 init() 一般用于初始化操作
//没有入参 没有返回值
func init() {
	fmt.Println("main init")
}

func main() {
	calc.Add(2, 3)
	fmt.Println("sss")
}
