package main

import (
	"bufio"
	"fmt"
	"os"
)

//获取用户输入如果有空格

func useScan() {
	var s string
	fmt.Print("请输入:")
	fmt.Scan(&s)
	fmt.Println("输入的内容是: ", s)

}

func useBufio() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	fmt.Println("输入的内容是: ", s)

}

func main() {

	//useScan()
	useBufio()
}
