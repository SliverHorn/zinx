package service

import (
	"errors"
	"fmt"
	"github.com/SliverHorn/zinx/interfaces"
	"net"
)

type Server struct {
	// 服务器监听的IP
	IP string

	// 服务器监听的端口
	Port int

	// 服务器的名称
	Name string

	// 路由
	Router interfaces.Router

	// 服务器绑定的IP版本
	IPVersion string
}

// NewServer Server的构造函数
func NewServer(IP string, port int, name string, IPVersion string) interfaces.Server {
	return &Server{IP: IP, Port: port, Name: name, Router:nil, IPVersion: IPVersion}
}

func (s *Server) Stop() {
	// TODO 将一些服务器的资源,状态或者一些开辟的链接信息,进行停职或者回收
	fmt.Println("stop")
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP: %v, Port: %d, is starting", s.IP, s.Port)
	go tcp(s)
}

func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()

	// TODO 做一些启动服务器之后的额外业务

	// 阻塞状态
	select {}
}

func (s *Server) AddRouter(router interfaces.Router)  {
	s.Router = router
	fmt.Println("添加路由成功")
}

func tcp(s *Server) {
	// 1. 获取一个TCP的Addr
	addr, rErr := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if rErr != nil {
		fmt.Println("func net.ResolveIPAddr() Failed! err:" + rErr.Error())
		return
	}

	// 2. 监听服务器的地址
	listenner, lErr := net.ListenTCP(s.IPVersion, addr)
	if lErr != nil {
		fmt.Println("func net.ListenTCP() Failed! err:" + lErr.Error())
	}
	fmt.Printf("Start %v Server success, Listenning\n", s.Name)

	var cid uint32
	cid = 0
	// 3.阻塞的等待客户端链接,处理客户端链接业务(读写)
	for {
		conn, err := listenner.AcceptTCP() // 如果有客户端链接过来,阻塞会返回
		if err != nil {
			fmt.Println("func listenner.AcceptTCP() Failed! err:" + err.Error())
			continue
		}
		// 将处理新链接的业务方法和conn进行绑定,得到我们的链接模块
		dealConn := NewConnection(conn, cid, s.Router)
		cid++
		// 启动当前的链接业务处理
		go dealConn.Start()
	}
}

// v0.1 调用 go echo(conn)
func echo(conn *net.TCPConn) {
	for {
		buf := make([]byte, 512)
		if count, err := conn.Read(buf); err != nil {
			fmt.Println("func conn.Read() Failed! err:" + err.Error())
			continue
		} else { // 回显
			fmt.Printf("接收到的信息为:%s, count:%d\n", buf, count)
			if _, err := conn.Write(buf[:count]); err != nil {
				fmt.Println("func conn.Write() Failed! err:" + err.Error())
				continue
			}
		}
	}
}

// v0.2 使用
func CallBackToClient(conn *net.TCPConn, content []byte, count int) error {
	if _, err := conn.Write(content[:count]); err != nil {
		fmt.Println("func conn.Write() Failed! err:", err.Error())
		return errors.New("CallBackToClient Error")
	}
	return nil
}