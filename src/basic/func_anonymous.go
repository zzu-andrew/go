package main

import "fmt"

func main() {
	/*
		 匿名函数：
			没有名字的函数
			匿名函数，需要在匿名函数尾部加上()直接调用
	*/
	func1()
	// 匿名函数采用的也是函数名字 ()实现调用
	// 匿名函数的实现和调用
	func() {
		fmt.Println("Anonymous call back.")
	}() // 加上了小括号实现了函数的调用

	// 匿名函数的多次调用
	func3 := func() {
		fmt.Println("func3 call func.")
	} // 不加返回值的时候直接返回对应的函数变量，供后期调用

	func3()
	// 带有参数的匿名函数调用， 需要在调用的括号中加上实参
	func(arg1, arg2 int) {
		fmt.Println(arg1, arg2)
	}(1, 2)


	// 带返回值的匿名函数
	res1 := func(a, b int) int {
		return a + b
	}(1, 2) // 除了增加小括号剩下的和定义的地方是一样的

	fmt.Println(res1)

	// 函数回调
	retValue := opener(20, 10, add)
	fmt.Println(retValue)

	fmt.Println("func anonymous demo")
}

func func1() {
	fmt.Println("I'm fun1() function.")
}

func add(a, b int) int {
	return a + b
}

func opener(a int, b int, fun func(int, int) int) int {
	fmt.Println("call func fun.")
	return fun(a, b)
}
