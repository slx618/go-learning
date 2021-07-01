package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	var rev [1024]byte
	for {
		msg, _ := reader.ReadString('\n')
		//发数据
		conn.Write([]byte(msg))

		//收数据
		n, _, err := conn.ReadFromUDP(rev[:])
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("client收到", string(rev[:n]))
	}
}
