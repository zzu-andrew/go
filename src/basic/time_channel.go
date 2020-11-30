package main

import (
	"fmt"
	"time"
)

func main() {

	/*
	1. func NewTimer(d Duration) *Timer
		创建一个定时器，d时间后触发
	*/

	timer := time.NewTimer(3*time.Second)
	fmt.Println("%T\n", timer)
	fmt.Println(time.Now())

	//此处等待 channel
	ch2 := timer.C
	fmt.Println(<- ch2)

	fmt.Println("channel timer")
}
