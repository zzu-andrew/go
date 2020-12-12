package main

import "fmt"

// 对已经关闭的通道进行取值，能够取到不过ok返回的false

func main() {
	ch1 := make(chan bool, 2)
	ch1 <- true
	ch1 <- true
	// 关闭通道
	close(ch1)

	// 通道即使关闭也是能从通道中取出剩余的值的
	<- ch1
	<- ch1

	x, ok := <- ch1
	fmt.Println(x, ok)
	x, ok = <- ch1
	fmt.Println(x, ok)
	x, ok = <- ch1
	fmt.Println(x, ok)
}
