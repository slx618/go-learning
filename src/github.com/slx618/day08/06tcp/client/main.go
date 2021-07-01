package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//1. 与 server 建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("建立连接失败", err)
	}
	//2. 发生数据
	var msg string
	//if len(os.Args) < 2 {
	//	msg = "hello world"
	//} else {
	//	msg = os.Args[1]
	//}
	reader := bufio.NewReader(os.Stdin)
	for {
		//fmt.Scan(&msg)
		msg, _ = reader.ReadString('\n')
		if msg == "exit" {
			break
		}
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("发生数据失败", err)
		}
	}
	conn.Close()
}
