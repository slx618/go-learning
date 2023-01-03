package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})

	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	// 不需要建立链接
	var data [1024]byte
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			n, addr, err := conn.ReadFromUDP(data[:])
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(string(data[:n]))
			revData := strings.ToUpper(string(data[:n]))
			//发送数据
			_, err = conn.WriteToUDP([]byte(revData), addr)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}()
	wg.Wait()
}
