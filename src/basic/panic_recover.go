package main

import "fmt"

func main() {
	/*
		panic recover
		panic让当当前程序停止运行，recover让程序回复运行，但是recover必须在defer函数中运行
	*/
	funcA()
	defer myprint("defer main:3.......")
	funcB()
	defer myprint("defer main:4.....")

	fmt.Println("main over")
}

func myprint(s string) {
	fmt.Println(s)
}

func funcA() {
	fmt.Println("这时函数funcA的测试函数")
}

func funcB() { // 有defer的函数是外围函数
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg, "程序恢复啦....")
		}
	}()

	fmt.Println("我是函数funcB()的测试函数...")
	defer myprint("defer funcB 1")
	for i := 1; i <= 10; i++ {

		fmt.Println("i:", i)
		if i == 5 {
			panic("funcB 恐慌....")
		}
	}
	defer myprint("defer funcB 2")
}
