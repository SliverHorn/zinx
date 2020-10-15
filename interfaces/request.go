package interfaces

// 实际上是把客户端请求的链接信息和请求的数据包装到一个Request中
type Request interface {
	// 得到请求的消息数据
	GetData() []byte

	// 得到当前的链接
	GetConnection() Connection
}