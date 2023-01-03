package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeDemo1() {
	//0000 0001 和 0000 0000 与后一直是 0000 0001
	handle, err := os.OpenFile("xx.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	//write
	n, err := handle.Write([]byte("xxxx\n"))
	fmt.Println(n)
	//write string
	handle.WriteString("什么意思\n")
	handle.Close()
}

func writeDemo2() {
	handle, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer handle.Close()
	w := bufio.NewWriter(handle)
	for i := 10; i > 0; i-- {
		w.WriteString("我是新的一行\n") // 写到缓存中
	}
	w.Flush()
}

func writeDemo3() {
	str := "xxxhewq租房处"
	err := ioutil.WriteFile("xxx.txt", []byte(str), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//打开文件写内容
func main() {
	writeDemo2()
	writeDemo3()
}
