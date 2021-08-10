//包名 main 包只能有一个
package main

//导入语句
import "fmt"

//函数外只能放置标识符(变量 常量 函数 类型)的声明

//CGO_ENABLED=0 GOOS=linux GOARCH=amd64
//CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
//CGO_ENABLED=0 GOOS=windows GOARCH=amd64
//go build -o xxx 输出文件
//go install 先编译再移动文件
//go build 是以 src 开始取找文件的
func main() {
	fmt.Println("hello world")
}
