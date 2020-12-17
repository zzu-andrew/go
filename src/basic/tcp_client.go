package main

import (
	"fmt"
	"net"
)

// tcp client
func main() {
	//	1. 建立链接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("dial failed.")
		return
	}
	defer func() {
		_ = conn.Close()
	}()

	//	2. 发送数据
	// 多次发送短数据， 粘包
	for i := 0; i < 10; i++ {
		_, err := conn.Write([]byte("hello world!"))
		if err != nil {
			fmt.Println("send failed")
		}
		//time.Sleep(time.Second)
	}
}
