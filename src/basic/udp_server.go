package main

import (
	"fmt"
	"net"
	"strings"
)

//UDP server
func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("listen UDP failed, err:", err)
		return
	}

	//	不需要建立链接 直接收发数据
	var data [1024]byte
	for {
		// data[:] 的意思就是将数组按照切片传入到一个函数中， 而且是引用的方式传入的
		// 利用了切片是对数组的引用的形式
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read data failed, err =", err)
			break
		}
		fmt.Println(string(data[:n]))
		replay := strings.ToUpper(string(data[:n]))
		// 发送数据
		n, err = conn.WriteToUDP([]byte(replay), addr)
		fmt.Println(n, err)
	}
}
