package interfaces

import "net"

// 定义链接模块的抽象层
type Connection interface {
	// 停止链接:结束当前链接的工作
	Stop()

	// 发送数据:将数据发送给客户端
	Send()

	// 启动链接:允许客户端链接服务端
	Start()

	// 获取获取链接模块的链接ID
	GetConnID() uint32

	// 获取远程客户端的TCP状态IP port
	RemoteAddress() net.Addr

	// 获取当前链接模块的绑定socket conn
	GetTCPConnection() *net.TCPConn

}


// 定义一个处理链接业务的方法
type HandleFunc func(conn *net.TCPConn, content []byte, count int) error