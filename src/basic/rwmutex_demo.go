package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁
var rwMutex *sync.RWMutex

// 等待组
var rwWg *sync.WaitGroup

func main() {
	rwMutex = new(sync.RWMutex)
	rwWg = new(sync.WaitGroup)

	rwWg.Add(3)
	//go readData1(1)
	//go readData1(2)
	go writeData1(1)
	go readData1(2)
	go writeData1(3)
	rwWg.Wait()

	fmt.Println("rwmutex demo")
}

/* 读数据操作能同时开始 */
func readData1(i int) {
	defer rwWg.Done()
	rwMutex.RLock()
	fmt.Println("read start......", i)
	fmt.Println("读数据......", i)
	time.Sleep(1 * time.Second)
	fmt.Println("读数据结束....", i)
	rwMutex.RUnlock()

}

func writeData1(i int) {
	defer rwWg.Done()
	rwMutex.Lock()
	fmt.Println("开始：write data.....", i)
	fmt.Println("正在写数据......", i)
	time.Sleep(1 * time.Second)
	fmt.Println("写结束..........", i)
	rwMutex.Unlock()
}
