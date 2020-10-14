package main

import "github.com/SliverHorn/zinx/service"

func main() {
	s := service.NewServer("0.0.0.0",9876, "[zinx V0.2]", "tcp4")
	s.Serve()
}