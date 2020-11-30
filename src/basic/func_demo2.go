package main

import (
	"fmt"
)

func main() {

	ret := getSumRet()
	fmt.Println(ret)
	peri, area := getRectangle(3, 5)
	fmt.Println(peri, area)
	//	 递归函数的调用
	//	 函数自己调用自己
	subSum := getSunSum(5)
	fmt.Println(subSum)

	/*
		1	2	3	4	5	6	7	8	9	10	11	12
		1	1	2	3	5	8	13	21	34	55	89	144
	*/
	fmt.Println("getFibonacci is :", getFibonacci(12))

	// 函数类型
	// 函数类型 关键字func替换 函数名之后，就是函数类型
	// 函数的类型 func(参数列表的类型)(函数返回值列表的类型)
	fmt.Printf("%T\n", func3) // 函数类型 func()
	fmt.Printf("%T\n", func4) // func() int
	fmt.Printf("%T\n", func5) // func(int, float32) (int, float32)
	// 函数名 和函数名+（）含义不一样   单纯的函数名代表的是一个指向函数的地址，使用()之后才是对函数的调用
	// 函数的调用，其实也是指向对应内存空间的 代码
	fmt.Println(func3) // 0x49b2c0 不同的计算机可能打印出来不是一样的

	var c func()
	fmt.Println(c)
	c = func3
	// 函数的调用
	c()
	// 函数在go语言中是一种复合类型，可以看做是特殊的变量



	//	函数的类型
}

// ctrl + shift + alt + 7 == 搜索函数调用的所有地方
func getSumRet() (ret int) {
	var retValue int
	fmt.Println(retValue)
	return retValue
}

// 函数的返回值能像参数一样操作， 其实go的实现也是按照这种原理进行实现的
// 将函数的返回值进行赋值之后，return 实际接收的时候是从栈中取值
func getRectangle(len, wid float64) (peri float64, area float64) {
	peri = (len + wid) * 2
	area = len * wid
	fmt.Println(peri, area)
	return
}

// 函数的递归调用
func getSunSum(num int) (sum int) {
	if num == 1 {
		return num
	}
	return getSunSum(num-1) + num
}

// 斐波那契数列
func getFibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return getFibonacci(n-1) + getFibonacci(n-2)
}

func func3() {
	fmt.Println("call func3")
}
func func4() int {
	return 0
}

func func5(arg1 int, arg2 float32) (int, float32) {

	return arg1, arg2
}
