package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// udp client
func main() {
	sock, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("链接服务器失败，err:", err)
		return
	}
	defer func() {
		_ = sock.Close()
	}()
	var replay [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("input something.")
		msg, _ := reader.ReadString('\n')

		n, _ := sock.Write([]byte(msg))
		fmt.Println(n)
		//	 收取回复的数据
		n, addr, err := sock.ReadFromUDP(replay[:])
		if err != nil {
			fmt.Println(addr, err)
		}

		fmt.Println(addr, n, string(replay[:n]))
	}

}
