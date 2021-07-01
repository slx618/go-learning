package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//大端小端 BigEndian LittleEndian  从左往右写 和 从右往左写
func main() {
	// 1. 本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("开启监听失败", err)
		return
	}
	// 2. 等待别人来跟我建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("建立连接失败", err)
			return
		}

		// 3. 与客户端通信
		go func() {
			defer conn.Close()
			var tmp [128]byte
			reader := bufio.NewReader(os.Stdin)

			for {
				n, err := conn.Read(tmp[:])
				if err != nil {
					fmt.Println("读取数据失败", err)
					return
				}
				fmt.Println(string(tmp[:n]))
				fmt.Println("回复:")
				msg, _ := reader.ReadString('\n')
				conn.Write([]byte(msg))

			}
		}()

	}
}
