package main

import (
	"fmt"
	"net"
	"time"
)

// 模拟客户端的请求
func main() {
	time.Sleep(1*time.Second)
	// 1. 链接server, 得到conn链接
	if conn, err := net.Dial("tcp", "127.0.0.1:9876"); err != nil {
		fmt.Println("链接失败")
		return
	} else {
		for {
			if _, err := conn.Write([]byte("SliverHorn Call Server")); err != nil {
				fmt.Println("Call failed, err:", err)
				return
			}
			buf := make([]byte, 512)
			if count, err := conn.Read(buf); err != nil{
				fmt.Println("Read failed, err:", err)
				return
			} else {
				fmt.Printf("Server Call Back:%s, count:%d\n", buf, count)
			}
			// CPU阻塞
			time.Sleep(7*time.Second)
		}
	}
}


