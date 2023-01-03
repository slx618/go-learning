package main

import (
	Proto "github.com/slx618/day08/09nianbao_solve"
	"net"
	"time"
)

//
func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:30000")
	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := "hello world xxx"
		b, _ := Proto.Encode(msg)
		conn.Write(b)
		time.Sleep(time.Second)
	}

}
