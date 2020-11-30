package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		关闭通道：close(ch)
		   子goroutine；写10个数据
		    每写一个，阻塞一次，主goroutine读取一次 解除阻塞
			主goroutine 读取数据
		      每次读取数据，阻塞一次， 子goroutine每次写入一个就解除主阻塞
	*/
	cha1 := make(chan int)
	go sendData1(cha1)
	//读取数据.使用for循环需要自己判断是否通道关闭
	/*	for {
			v, ok := <-cha1
			if !ok {
				fmt.Println("已经读取了所有的数据......", ok)
				break
			}
			fmt.Println("读取的数据....", v)
		}
	*/
	// 使用for range就可以让range自己判断通道啥时候关闭
	for v := range cha1 {
		fmt.Println("main ..over....", v)
	}
	//单向 只能写数据，不能读数据
	//ch1 := make(chan <- int)
	// 单向只能读不能写数
	//ch2 := make(<- chan int)

	fmt.Println("channel close")
}

func sendData1(cha1 chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1000 * time.Millisecond)
		cha1 <- i
	}
	close(cha1)
}
