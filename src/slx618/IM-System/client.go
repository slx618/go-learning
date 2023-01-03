package main

import (
	"fmt"
	"net"
	"flag"
	"io"
	"os"
)


type Client struct {
	ServerIp string
	ServerPort int
	Name string
	conn net.Conn
	flag int
}

func NewClient(ip string, port int) *Client {

	client := &Client{
		ServerIp: ip,
		ServerPort: port,
		flag: 999,
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println("connect error:", err)
		return nil
	}

	client.conn = conn

	return client
}

func (c *Client) menu() bool {
	fmt.Println("1. 公聊模式")
	fmt.Println("2. 私聊模式")
	fmt.Println("3. 修改用户名")
	fmt.Println("0.  退出")

	var flag int
	fmt.Scanln(&flag)

	fmt.Println("flag", flag)

	if flag >=0 && flag <= 3 {
		c.flag = flag
		return true
	} else {
		fmt.Println(">>>请输入合法的数字")
		return false
	}

}

func (c *Client) Rename() bool {
	fmt.Println(">>>请输入用户名")
	fmt.Scanln(&c.Name)
	fmt.Println(c.Name)

	_, err := c.conn.Write([]byte("rename|" + c.Name))
	if err != nil {
		fmt.Println("conn Write error: ", err)
		return false
	}

	return true

}


func (c *Client) PublicChat() {
	var chatMsg string
	

	for {
		fmt.Println(">>>请输入聊天内容, exit 退出")
		fmt.Scanln(&chatMsg)
		if chatMsg == "exit" {
			break
		} 

		if len(chatMsg) > 0 {
			_, err := c.conn.Write([]byte(chatMsg))
			if err != nil {
				fmt.Println("PublicChat error: ", err)

			}
		}
	
		chatMsg = ""
	}

}

func (c *Client) who() {
	fmt.Println(">>>在线用户列表")

	_, err := c.conn.Write([]byte("who\n"))
	if err != nil {
		fmt.Println("who error: ", err)
	}

}

func (c *Client) To() bool {
	var to string
	var msg string

	for {
		c.who()

		fmt.Println(">>>请输入私聊,聊天对象的用户名 exit 退出")
		fmt.Scanln(&to)
		if to == "exit" {
			break
		}
		
		for {

			fmt.Println(">>> 请输入消息内容, exit 退出")
			fmt.Scanln(&msg)
			if msg == "exit" {
				break
			}
			if len(msg) > 0 {
				_, err := c.conn.Write([]byte("to|"+ to + "|" + msg))
				if err != nil {
					fmt.Println("PublicChat error: ", err)
				}
			}
			msg = ""
	
		}

		to = ""
		
	}

	_, err := c.conn.Write([]byte("to|" + c.Name))
	if err != nil {
		fmt.Println("conn Write error: ", err)
		return false
	}

	return true

}

func (c *Client) dealResponse() {
	io.Copy(os.Stdout, c.conn)

}

func (c *Client) Run() {
	for c.flag != 0 {
		// 一直出这个菜单
		for c.menu() != true {}

		switch c.flag {
			
			case 1:
				c.PublicChat()
				break
			case 2:
				c.To()				
				break
			case 3:
				c.Rename();

				break
		}

	}
}

var serverIp string
var serverPort int
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器 IP地址(默认 127.0.0.1)")
	flag.IntVar(&serverPort, "port", 9970, "设置服务器 端口(默认 9970)")
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)

	if client == nil {
		fmt.Println("连接服务器失败...")
	}

	go client.dealResponse()

	fmt.Println("连接服务器成功...")
	client.Run()
	// select {

	// }

}