package main

import (
	"fmt"
	"net"
)

func main() {
	//1. 与 server 建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("建立连接失败", err)
	}
	//2. 发生数据
	_, err = conn.Write([]byte("hello world"))
	if err != nil {
		fmt.Println("发生数据失败", err)
	}
	conn.Close()
}
