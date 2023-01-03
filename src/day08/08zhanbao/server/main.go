package main

import (
	"bufio"
	"fmt"
	"github.com/slx618/day08/09nianbao_solve"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	//var buf [1024]byte
	for {
		//n, err := reader.Read(buf[:])
		revStr, err := Proto.Decode(reader)

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("接受到的数据:", revStr)
	}

}

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:30000")
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go process(conn)
	}
}
