package main

import "fmt"

func main() {
	/*
		  defer的词义： 延迟 推迟
		在go语言中，使用defer关键字来延迟一个函数或者方法的执行
		1. defer 函数或方法 一个函数或方法会被延迟执行
		2. defer的使用
			A 对象文件的close 临时文件的删除
			B go语言中关于异常的处理，使用panic 和recover
				panic函数引发的恐慌 导致程序中断执行
				recover函数用于恢复程序的执行recover()语法上要求必须在defer中执行
		3. 如果多个defer函数
			倒序调用 最后defer调用的最先被调用


	*/

	defer testDeferPrint("defer test1")
	defer testDeferPrint("defer test2")
	fmt.Println("这里先打印，最后执行defer关键字后面调用的函数")

	//fmt.Println("defer demo")
}

func testDeferPrint(testStr string) {
	fmt.Println(testStr)
}
