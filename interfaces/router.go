package interfaces

// 路由抽象接口:路由里的数据都是Router
type Router interface {
	// 处理conn业务之前的钩子方法Hook
	BeforeHandle(request Request)
	// 处理conn业务主方法Hook
	NowHandle(request Request)
	// 处理conn业务之后的钩子方法Hook
	AfterHandle(request Request)
}