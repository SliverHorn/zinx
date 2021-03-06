package interfaces

// 定义一个服务器接口
type Server interface {
	// 停止服务器
	Stop()

	// 启动服务器
	Start()

	// 运行服务器
    Serve()

	// 路由功能
	AddRouter(router Router)
}
