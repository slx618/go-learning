package chapter1

import "fmt"

var i, j int = 1, 23
var b, n, m = true, 132, "haha"
func main()  {
	//var z int = 123 函数内部的变量需要被调用 否则报错
	//h := 23
	fmt.Println(i, j)
	fmt.Println(b, n, m)
}
