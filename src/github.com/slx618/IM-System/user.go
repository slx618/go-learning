package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
	server *Server
}

 
func NewUser(conn net.Conn, s *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C: make(chan string),
		conn: conn,
		server: s,
	}

	go user.ListenMsg()

	return user
}

func (u *User) ListenMsg() {
	for {
		msg := <-u.C

		u.conn.Write([]byte(msg + "\n"))
	}
}

func (u *User) Online() {
	//用户上线
	u.server.mapLock.Lock()
	u.server.OnlineUsers[u.Name] = u
	u.server.mapLock.Unlock()

	u.server.Broadcast(u, "用户上线")
}

func (u *User) Offline() {
	//用户下线
	u.server.mapLock.Lock()
	delete(u.server.OnlineUsers, u.Name)
	u.server.mapLock.Unlock()

	u.server.Broadcast(u, "用户下线")
}

func (u *User) DoMsg(msg string) {
	if msg == "who" {
		//查询当前用户列表
		u.server.mapLock.Lock()
		for _, user := range u.server.OnlineUsers {
			userMsg := "[" + user.Addr + "]" + " " + user.Name + ": " + "在线...";
			u.ms(userMsg)
		}
		u.server.mapLock.Unlock()
	} else if len(msg) > 4 && msg[:3] == "to|" {
		to := strings.Split(msg, "|")[1]
		content := strings.Split(msg, "|")[2]
		if content == "" {
			u.ms("消息不能为空")
			return
		}

		if to == "to" {
			u.ms("消息格式不正确, 请使用 to|张三|消息 格式")
			return
		}
		u.dd(to, content)

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		nickname := msg[7:]

		//判断是否存在改昵称
		_, ok := u.server.OnlineUsers[nickname]
		if ok {
			u.ms("昵称已被占用: " + nickname)
		} else {
			u.server.mapLock.Lock()
			delete(u.server.OnlineUsers, u.Name)
			u.Name = nickname
			u.server.OnlineUsers[nickname] = u
			u.server.mapLock.Unlock()
			u.ms("设置昵称成功" + nickname)
		}

	} else {
		
		if len(msg) > 0 {
			u.server.Broadcast(u, msg)
		} else {
			u.ms("消息不能为空")
			return
		}
	}
}


func (u *User) ms(msg string) {
	u.conn.Write([]byte(msg + "\n"))
}

func (u *User) dd(nickname, content string) {
	remoteUser, ok := u.server.OnlineUsers[nickname]
	if !ok {
		u.ms(nickname + "用户不存在")
		return
	}
	remoteUser.conn.Write([]byte(u.Name + "私信你:" + content + "\n"))
}