package main

import (
	"fmt"
	"runtime"
)

func main() {
	//	获取goroot
	fmt.Println("GoROOT -->", runtime.GOROOT())
	/*获取操作系统， 是个常量*/
	/*
		GoROOT --> /usr/local/go
		os platform--< linux
	*/
	fmt.Println("os platform--<", runtime.GOOS)
	/*获取逻辑cpu的数量*/
	fmt.Println("逻辑CPU的数量-->", runtime.NumCPU())

	// gosched
	go func() {
		for i := 0; i < 35; i++ {
			fmt.Println("goroutine ....", i)
		}
	}()

	for i := 0; i < 36; i++ {
		// 让出时间片 让别的线程先执行
		runtime.Gosched()
		fmt.Println("main .......", i)
	}
}
