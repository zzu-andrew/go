package main

import "fmt"

func main() {
	// goy语言支持函数式编程：
	//	支持将一个函数的返回值作为另外一个函数的参数
	// 也支持将一个函数作为一个函数的返回值

	/**
	闭包(closure)：

	*/
	res1 := increment()
	fmt.Printf("%T\n", res1)
	// 局部变量应该销毁，但是还是能正常返回
	v1 := res1()
	fmt.Println(v1)
	fmt.Println("func closure")
	v2 := res1()
	fmt.Println(v2)
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	//	使用返回函数再次调用
	// 每当调用 increment的时候
	fmt.Println("call again")
	res2 := increment()
	v3 := res2()
	fmt.Println(v3)
}

/*
1. 如果将一个内层函数作为函数返回值
2. 内层函数中又涉及到外层函数的局部变量(自己定义或者外部传进来的参数，都是局部变量)
3. 1+2 条件满足后，就会导致该局部变量的声明周期发生改变，外层函数的局部变量不会随着外层函数的借宿而销毁
4. 这种内层函数变量和外层局部变量，统称为闭包结构
5. 闭包结构中，局部变量的声明周期会发生改变
*/
func increment() func() int { // 外层函数
	// 定义一个局部变量
	i := 0
	// 定义一个匿名函数， 并且进行返回
	fun := func() int { // 内层函数
		i++
		return i
	}
	// 返回该匿名函数
	return fun
}
