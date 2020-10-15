package service

import "github.com/SliverHorn/zinx/interfaces"

// 提供基类
// 提供需要继承的基类 基类方法为空
// 好处是如果不需要对应的hook 也不需要强行实现

// 实现Router时,先嵌入这个BaseRouter基类,然后根据需要对这个基类的方法进行重写即可
type BaseRouter struct{}

// 处理conn业务之前的钩子方法Hook
func (b *BaseRouter) BeforeHandle(request interfaces.Request) {}

// 处理conn业务主方法Hook
func (b *BaseRouter) NowHandle(request interfaces.Request)    {}

// 处理conn业务之后的钩子方法Hook
func (b *BaseRouter) AfterHandle(request interfaces.Request)  {}
