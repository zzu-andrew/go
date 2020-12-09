package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		一个goroutine 打印数字，另外一个goroutine打印字母，观察运行结果
	*/

	// 先创建子goroutine 执行printNum
	go printNum()
	//	main中打印字母
	for i := 0; i < 100; i++ {
		fmt.Println("\t主goroutine 打印字母。。。。。。。。。")
	}
	time.Sleep(1 * time.Second)
}

func printNum() {
	for i := 0; i < 1000; i++ {
		fmt.Println("子goroutine 打印数字：", i)
	}
}
