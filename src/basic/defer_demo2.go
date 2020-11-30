package main

import "fmt"

func main() {
	/*
		4. defer函数传参数的时候
			defer 函数在调用的时候已经传递参数了，只是代码没有执行而已

		5. defer函数注意点
	*/
	a := 2
	fmt.Println(a)
	defer func2(a)
	a++
	fmt.Println(a)


}

//
func func2(a int) {
	fmt.Println("defer 打印 a = ", a)
}

