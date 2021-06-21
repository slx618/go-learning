package calc

import "fmt"

func init() {
	fmt.Println("calc init")
}

//包中的变量名 函数名 接口名 结构体 如果是小写 表示私有 只能在本包中使用
//首字母大写的标识符 可以被外部调用
func Add(x, y int) int {
	fmt.Println("add")
	return x + y
}
