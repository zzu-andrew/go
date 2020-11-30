package main

import "fmt"

func main() {
/*
	num := 19
	num--
	if num > 10 {
		fmt.Println("num is grate than 10")
	} else {
		fmt.Println("num is low than 10")
	}
*/
	// if语句的其他写法
	/*
		statement用于声明变量，这些变量只能在当前的if语句中使用
		这个是其他语言不支持的，只有go语言中支持这种语法
		if statement; condition{

		}
	*/

	if num := 10; num % 2 == 0 {
		fmt.Println("num is even", num)
	} else {
		fmt.Println("num is odd", num)
	}

}
