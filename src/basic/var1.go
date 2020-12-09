package main

import "fmt"

func main() {
	//变量测试
	//小块内存 可以改变，不可以改变的是常量
	// var 变量名 类型
	var num1 int
	num1 = 30
	fmt.Printf("num1 = %d\n", num1)

	var num2 int = 15
	fmt.Printf("num2 = %d\n", num2)
	// 虽然go语言是静态余语言但是支持类型的自动判断
	var num3 = 30
	fmt.Println(
		"num3 = %d",
		num3,
	)
	fmt.Print("var test")
	// 使用类型腿短 := ,这种方式只能用于函数体内，不能用于全局变量的盛行声明
	num4 := 45
	fmt.Println("type = [%T]num4 = [%d]", num4, num4)

	// 变量集合声明
	var (
		studentName = "xiaoming"
		teacherName = "jianguo"
		studentAge  = 18
		teacherAge  = 28
	)
	fmt.Println("student name = [%s] age = [%d], Teacher name = [%s] age = [%d]",
		studentName, studentAge, teacherName, teacherAge)

}
