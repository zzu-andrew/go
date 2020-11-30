package main

import "fmt"

func main() {

	var a int8 = 8
	var b int16
	b = int16(a)

	fmt.Println(a, b)

	f1 := 16.86
	var c int
	c = int(f1)
	fmt.Println(f1, c)
	// 不是所有类型都能装换，bool类型不支持强制转换
	//cannot convert b1 (type bool) to type int8
	//b1 := true
	//a = int8(b1)
	//fmt.Println(a)
}
