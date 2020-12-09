package main

import "fmt"

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Print(x + y)
}

func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		fmt.Println("hello")
		f(x, y)
	}

	return tmp
}

func main() {
	//基本的数据类型
	//	1.布尔类型
	//	2.数值类型

	//	3.字符串
	//	4.复合数据类型
	//	 array slice mao function pointer struct interface channel...
	//	bool
	var a bool = true
	fmt.Printf("%t\n", a)

	//	byte == uint8
	//	rune == int32
	var b rune = 100
	fmt.Println(b)
	// 匿名函数，立即执行函数
	func(x, y int) {
		fmt.Println(x + y)
	}(1, 2)

	//	闭包
	f1(f3(f2, 2, 3))

}
