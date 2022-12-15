package main

import (
	"fmt"
	"net"
)


type Client struct {
	ServerIp string
	ServerPort int
	Name string
	conn net.Conn
}

func NewClient(ip string, port int) *Client {

	client := &Client{
		ServerIp: ip,
		ServerPort: port,
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("connect error:", err)
		return nil
	}

	client.conn = conn

	return client
}

func main() {
	client := NewClient("127.0.0.1", 9970)

	if client == nil {
		fmt.Println("连接服务器失败...")
	}



}