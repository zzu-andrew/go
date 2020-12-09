package main

import "fmt"

/*
闭包是什么？
闭包是一个函数，这个函数包含了他外部作用域的变量

函数可以作为返回值
函数内部查找变量的顺序，先自己内部找，找不到再往外层找

*/

/*
每次调用 increment都会重新生成一个闭包结构中的变量
内层函数+外层函数局部变量(可以是入参或者定义的局部变量) = 闭包结构

如果将一个内层函数作为函数返回值
内层函数中又涉及到外层函数的局部变量(自己定义或者外部传进来的参数，都是局部变量)
1+2 条件满足后，就会导致该局部变量的声明周期发生改变，外层函数的局部变量不会随着外层函数的借宿而销毁
这种内层函数变量和外层局部变量，统称为闭包结构
闭包结构中，局部变量的声明周期会发生改变
*/

// 老的函数接口，只能接受没有参数
func oldFun(f func()) {
	fmt.Println("this is f1")
	f()
}

// 新的接口
func newFun(x, y int) {
	fmt.Println("this is f2")
	fmt.Print(x + y)
}

//使用闭包接口可以实现将新接口，转换为老接口
//使用特性   闭包结构
func closureFun(f func(int, int), x, y int) func() {
	// 闭包结构中， x y 变量的生命周期已经发生改变，在任何调用匿名函数的地方
	//都可以得到 x, y 的值，从而实现，在oldFun函数调用匿名函数的时候，压栈，调到函数newFun的时候，依旧能获取到 x y的值
	fmt.Println("closure func")
	tmp := func() {
		fmt.Println("closure tmp func")
		f(x, y)
	}

	return tmp
}

func main() {

	// 匿名函数，立即执行函数
	func(x, y int) {
		fmt.Println(x + y)
	}(1, 2)

	//	闭包
	fmt.Println("-------------closure func--------------")
	fakeOldFun := closureFun(newFun, 2, 3)
	oldFun(fakeOldFun)

}
