package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	 if switch select分支选择语句
	*/

	cha1 := make(chan int)
	cha2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		cha1 <- 1
	}()

	go func() {
		time.Sleep(3 * time.Second)
		cha2 <- 20
	}()

	select {
	case num1 := <-cha1:
		fmt.Println(num1)
	case num2, ok := <-cha2:
		if ok{
			fmt.Println("num2 = ", num2)
		}else {
			fmt.Println("cha2 通道已关闭.....")
		}
	default:
		fmt.Println("default")
	}



	fmt.Println("select demo")
}
