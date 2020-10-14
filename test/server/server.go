package main

import "github.com/SliverHorn/zinx/service"

func main() {
	s := service.NewServer("0.0.0.0",9876, "[Server V0.1]", "tcp4")
	s.Serve()
}