package main

import (
	"fmt"
	"github.com/SliverHorn/zinx/interfaces"
	"github.com/SliverHorn/zinx/service"
)

type PingRouter struct {
    service.BaseRouter
}

func (b *PingRouter) BeforeHandle(request interfaces.Request) {
	fmt.Println("Call Router BeforeHandle...")
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("Before Ping")); err != nil {
		fmt.Println("Before Ping Error=", err.Error())
	}
}
func (b *PingRouter) NowHandle(request interfaces.Request) {
	fmt.Println("Call Router NowHandle...")
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("Ping-->Success")); err != nil {
		fmt.Println("Now Ping Error=", err.Error())
	}
}
func (b *PingRouter) AfterHandle(request interfaces.Request) {
	fmt.Println("Call Router AfterHandle...")
	if _, err := request.GetConnection().GetTCPConnection().Write([]byte("After Ping")); err != nil {
		fmt.Println("After Ping Error=", err.Error())
	}
}



func main() {
	s := service.NewServer("0.0.0.0",9876, "[zinx V0.3]", "tcp4")
	s.AddRouter(&PingRouter{})
	s.Serve()
}