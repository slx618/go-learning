package main

import "fmt"

func funcA() {
	fmt.Println("a")
}

func funcB() {

	//defer 要写在 panic前面
	defer func() {
		//recover 必须搭配 defer 使用
		err := recover()
		fmt.Println(err)
		fmt.Println("释放资源")
	}()
	panic("严重错误发生 ") //程序崩溃退出
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()

}
