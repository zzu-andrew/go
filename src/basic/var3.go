package main

import "fmt"

func main() {
	// iota 特殊常量，可以被编译器自动修改的常量，每当定义一个const， iota会自动清零
	// 在同一个变量集合中，iota随着变量的定义一直递增，直到遇到下一个const在清零
	const(
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	const(
		e = iota
		f    // 默认和上一行一致
		g = iota
	)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
//	枚举实现
	const(
		MALE = iota
		FEMALE
		UNKNOWM
	)
	fmt.Println(MALE, FEMALE, UNKNOWM)



}
