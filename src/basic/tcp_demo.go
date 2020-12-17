package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func processConn(conn net.Conn) {
	defer func() {
		_ = conn.Close()
	}()

	//	3. 与客户端通信
	reader := bufio.NewReader(conn)

	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据：", recvStr)
	}

}

// tcp server
func main() {
	//	1. 本地端口建立服务
	lister, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("start server failed.")
		return
	}
	defer func() {
		_ = lister.Close()
	}()

	for {
		//	2. 等待链接
		conn, err := lister.Accept()
		if err != nil {
			fmt.Println("accept failed, err = ", err)
			return
		}

		go processConn(conn)
	}

}
