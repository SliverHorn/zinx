package snet

import (
	"fmt"
	"github.com/SliverHorn/zinx/siface"
	"net"
)

type Server struct {
	// 服务器监听的IP
	IP string
	// 服务器监听的端口
	Port int
	// 服务器的名称
	Name string
	// 服务器绑定的IP版本
	IPVersion string
}

// NewServer Server的构造函数
func NewServer(IP string, port int, name string, IPVersion string) siface.IServer {
	return &Server{IP: IP, Port: port, Name: name, IPVersion: IPVersion}
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
	select {

	}
}

func tcp(s *Server) {
	// 1. 获取一个TCP的Addr
	addr, rErr := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if rErr != nil {
		fmt.Println("func net.ResolveIPAddr() Failed! err:"+rErr.Error())
		return
	}

	// 2. 监听服务器的地址
	listenner, lErr := net.ListenTCP(s.IPVersion, addr)
	if lErr != nil {
		fmt.Println("func net.ListenTCP() Failed! err:"+lErr.Error())
	}
	fmt.Printf("Start %v Server success, Listenning", s.Name)

	// 3.阻塞的等待客户端链接,处理客户端链接业务(读写)
	for {
		conn, err := listenner.AcceptTCP() // 如果有客户端链接过来,阻塞会返回
		if err != nil {
			fmt.Println("func listenner.AcceptTCP() Failed! err:"+err.Error())
			continue
		}
		go echo(conn)
	}
}

func echo(conn *net.TCPConn) {
	for  {
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("func conn.Read() Failed! err:"+err.Error())
			continue
		}

		// 回显
		if _, err := conn.Write(buf[:cnt]); err != nil {
			fmt.Println("func conn.Write() Failed! err:"+err.Error())
			continue
		}
	}
}