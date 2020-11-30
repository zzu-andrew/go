package utils

import "fmt"

func Count() {
	fmt.Print("utils package")
}
/*
同一个go文件中多个init函数， 谁先先的谁先被调用，后写的后被调用
若是一个pkg下有多个文件，会按照文件名字符串排序进行执行init函数
*/
func init() {
	fmt.Println("inti utils 1")
}

func init() {
	fmt.Println("inti utils 2")
}