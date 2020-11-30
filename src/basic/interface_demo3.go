package main

import "fmt"

func main() {
	/*
		接口的嵌套
	*/
	var cat Cat12 = Cat12{}
	cat.test1()
	cat.test2()
	cat.test3()

	fmt.Println("----------------------")
	// 因为当前a1就是AA1类型所以只能调用test1方法
	var a1 AA1 = cat
	a1.test1()

	var a2 BB1 = cat
	a2.test2()

	var a3 CC1 = cat
	a3.test1()
	a3.test2()
	a3.test3()


}

type AA1 interface {
	test1()
}

type BB1 interface {
	test2()
}

type CC1 interface {
	AA1
	BB1
	test3()
}

// 如果想实现接口C那么接口A和接口B中的方法也要实现
type Cat12 struct {
}

func (c Cat12) test1() {
	fmt.Println("test1.....")
}

func (c Cat12) test2()  {
	fmt.Println("test2.....")
}

func (c Cat12) test3()  {
	fmt.Println("test3......")
}
