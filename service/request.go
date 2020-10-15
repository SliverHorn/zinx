package service

import "github.com/SliverHorn/zinx/interfaces"

type Request struct {
	// 客户端请求的数据
	data []byte

	// 已经和客户端建立好的链接
    conn interfaces.Connection
}

func (r *Request) GetData() []byte {
	return r.data
}

func (r *Request) GetConnection() interfaces.Connection {
	return r.conn
}
