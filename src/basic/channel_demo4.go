package main

import (
	"fmt"
	"sync"
)

// 单项通道，就是在chan关键字上增加方向剪头的通道
var wg2 sync.WaitGroup

func wgFun1(ch1 chan<- int) {
	defer wg2.Done()
	for i := 1; i <= 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func wgFun2(ch2 chan<- int, ch1 <-chan int) {
	defer wg2.Done()

	for {
		x, ok := <-ch1
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
	wg2.Add(2)
	go wgFun1(ch1)
	go wgFun2(ch2, ch1)
	wg2.Wait()
	for v := range ch2 {
		fmt.Println("get ch2 is ", v)
	}

}
