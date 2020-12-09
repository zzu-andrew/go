package main

import (
	"fmt"
	math_rand "math/rand"
	"sync"
	"time"
)

/*
全局变量表示
*/
var ticket int = 10
var mutex sync.Mutex

func main() {
	/*
		4个 goroutine 同时卖票
	*/
	go saleTickets("1")
	go saleTickets("2")
	go saleTickets("3")
	go saleTickets("4")

	time.Sleep(20 * time.Second)
}

func saleTickets(name string) {
	for {
		//上锁
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(math_rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出", ticket)
			ticket--
		} else {
			mutex.Unlock()
			fmt.Println(name, "售完了，没有票了....")
			break
		}
		mutex.Unlock()

	}
}
