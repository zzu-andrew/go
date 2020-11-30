package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup // 创建同步等待组
func main() {

	wg.Add(2)
	go synFunc1()
	go synFunc2()
	// 表示主goroutine进入等待
	wg.Wait()

	fmt.Println("syn demo")
}

func synFunc1() {
	for i := 1; i < 10; i++ {
		fmt.Println("func 函数中打印...A", i)
	}
	wg.Done()
}

func synFunc2() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("func2 函数中打印...B", i)
	}

}
