package main

import (
	"fmt"
	"sync"
)

/*
goroutine与线程
可增长的栈
OS线程(操作系统的线程)一般有固定的栈内存，通常为2MB,一个goroutine的栈在其生命周期开始时只有很小的栈
典型情况下占用的事2KB,goroutine的栈一般不是固定的，可可以按需进行增大和缩小，goroutine的栈大小限制可达到1GB
虽然极少会用这么大，所以go语言一把创建十万个左右的goroutine也是可以的

*/

// 使用 @sync.WaitGroup 优雅的实现多个goroutine同步，而不是使用sleep进行
var groutinWg sync.WaitGroup

func goroutineHello(i int) {

	fmt.Println(i)
	groutinWg.Done()
}



func main() {

	for i := 0; i < 1000; i ++ {
		groutinWg.Add(1)
		go goroutineHello(i)
	}

	groutinWg.Wait()

}
