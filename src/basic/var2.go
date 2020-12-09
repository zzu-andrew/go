package main

import "fmt"

func main() {
	var num int
	num = 100
	fmt.Printf("num = %d &"+
		" = %p\n", num, &num)

	//	常量的定义
	//	常量尽量都使用大写的定义，定义之后不允许修改
	const PI = 3.14
	const PATH string = "http:www.baidu.com"
	fmt.Println(PATH)
	//	定义一组常量
	//	变量的集合
	const C1, C2, C3 = 100, 3, "haha"
	const (
		NAME     = "xiaoming"
		LINKMODE = 1
	)
	// 一组常量，如果某个常量没有给初始值，就会默认和上一个常量保持一致的值
	const (
		a int = 100
		b
		c string = "ruby"
		d
		e
	)
	fmt.Printf("type = %T, b = [%d]\n", a, a)
	fmt.Printf("type = %T, b = %d\n", b, b)
	fmt.Printf("type = %T, b = %s\n", d, d)
	fmt.Printf("type = %T, b = %s\n", e, e)
}
