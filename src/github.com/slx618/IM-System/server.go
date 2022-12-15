package main

import (
	"fmt"
	"net"
	"sync"
	"io"
	"time"
)

type Server struct {
	Ip   string
	Port int

	//online user map
	OnlineUsers map[string]*User

	mapLock sync.RWMutex

	// 广播 channel
	ServerChan chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:          ip,
		Port:        port,
		OnlineUsers: make(map[string]*User),
		ServerChan:  make(chan string),
	}
	return server
}

//run server

func (s *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("socket create fail: ", err)
		return
	}
	//close socket listen
	defer listener.Close()

	fmt.Println("socket listen on: ", s.Port)

	//启动广播监听
	go s.ListenServerChan()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("socket accept fail: ", err)
			continue
		}

		// do handle
		go s.Handle(conn)
	}

}

func (s *Server) Handle(conn net.Conn) {
	fmt.Println("socket handdle...")
	user := NewUser(conn, s)
	//用户上线 广播用户
	user.Online()

	isLive := make(chan bool)

	//接受客户端消息
	go func() {
		buff := make([]byte, 4096)
		for {
			n, err := conn.Read(buff)
			if n == 0 {
				// 客户端关闭了
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("conn read err", err)
				return
			}

			// 提前用户信息 去除\n
			msg := string(buff[:n-1])
			fmt.Println("消息", msg)
			//用户消息处理 广播
			user.DoMsg(msg)

			// 用户活跃
			isLive <- true
		}
	}()


	for {
		select {
			// 当前用户活跃
		case <-isLive :
			// 不用做任何操作
		
			//超时 600 秒触发
		case <-time.After(time.Second * 600):
			user.ms("超时不活跃")
			delete(s.OnlineUsers, user.Name)
			// 关闭管道
			close(user.C)

			//关闭连接
			conn.Close()

			// 退出 gorotinue
			return
		}
	}

}


func (s *Server) ListenServerChan() {
	for { 
		msg := <-s.ServerChan

		s.mapLock.Lock()
		for _, user := range s.OnlineUsers {
			user.C <-msg
		}
		s.mapLock.Unlock()
	}
}


func (s *Server) Broadcast(user *User, msg string) {
	senMsg := "[" + user.Addr + "] " + user.Name + ": " + msg + "\n";

	s.ServerChan <- senMsg
}
