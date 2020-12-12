package main

import (
	"fmt"
	"sync"
)

// 1. 创建goroutine1, 向 ch1中发送1,100个数
// 2. 创建goroutine2, 向 ch2中发送, 1，-1000的平方
// 3. main函数中从ch2将值取出

var wg1 sync.WaitGroup

func wg1Fun(ch1 chan int) {
	defer wg1.Done()
	for i := 1; i <= 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func wg2Fun(ch2 chan int, ch1 chan int) {
	defer wg1.Done()

	for {
		x, ok := <- ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch2)
}


func main() {
	var ch1 = make(chan int, 100)
	var ch2 = make(chan int, 100)
	wg1.Add(2)
	go wg1Fun(ch1)
	go wg2Fun(ch2, ch1)
	wg1.Wait()
	for v := range ch2 {
		fmt.Println("get ch2 is ", v)
	}


}
