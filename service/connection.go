package service

import (
	"fmt"
	"github.com/SliverHorn/zinx/interfaces"
	"net"
)

type Connection struct {
	// 当前链接的socket TCP套接字
	Conn *net.TCPConn

	// 链接ID
	ConnID uint32

	// 路由
	Router interfaces.Router

	// 告知当前链接已经退出/停止 channel
	ExitChan chan bool

	// 当前链接的状态
	isClosed bool

	// 当前链接锁绑定的处理业务方法API
	// handleAPI interfaces.HandleFunc
}

func NewConnection(conn *net.TCPConn, connID uint32, router interfaces.Router) *Connection {
	return &Connection{Conn: conn, ConnID: connID, isClosed: false, Router: router, ExitChan: make(chan bool, 1)}
}

func (c *Connection) Stop() {
	fmt.Println("关闭链接ing...ConnID=", c.ConnID)
	if c.isClosed {
		return
	}
	c.isClosed = true

	// 关闭socket链接
	defer c.Conn.Close()

	// 关闭资源
	close(c.ExitChan)
}

func (c *Connection) Send() {
	panic("implement me")
}

func (c *Connection) Start() {
	fmt.Println("开启链接ing...ConnID=", c.ConnID)
	go c.StartReader()
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddress() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// StartReader 链接的读业务
func (c *Connection) StartReader() {
	defer fmt.Printf("ConnID=%v, Remote Address=%v, 已退出\n", c.ConnID, c.RemoteAddress().String())
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中,最大512字节
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("读取失败,err", err.Error())
			continue
		}
		request := Request{data: buf, conn: c}
		go func(r interfaces.Request) {
			c.Router.BeforeHandle(r)
			c.Router.NowHandle(r)
			c.Router.AfterHandle(r)
		}(&request)

		//if err = c.handleAPI(c.Conn,buf,count); err != nil {
		//	fmt.Printf("ConnID=%v,err=%v", c.ConnID, err.Error())
		//	break
		//}
	}
}
