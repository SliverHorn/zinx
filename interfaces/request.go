package interfaces

type Request interface {
	// 得到请求的消息数据
	GetData() []byte

	// 得到当前的链接
	GetConnection() Connection
}