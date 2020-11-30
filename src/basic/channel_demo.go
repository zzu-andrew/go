package main

import "fmt"

func main() {
	/*
		通道名 channel
	*/
	/*刚创建的通道是nil*/
	var a chan bool
	fmt.Printf("%T,%v\n", a, a)
	a = make(chan bool)

	go func() {
		for i := 1; i < 10; i++ {
			fmt.Println("子goroutine is running....", i)
		}
		a <- true
		fmt.Println("结束....")
	}()
	fmt.Println("main....over.....")

	// 读取的操作是注阻塞的
	data := <-a
	fmt.Println(data)

	fmt.Println("channel demo")
}
