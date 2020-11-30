package main

import "fmt"

func main() {
	/*
		空接口： interface{}
		 不包含任何方法的接口， 正因为如此，所有的类型都实现了空接口，因此空接口可以存储任意类型的数值
	fmt print协议函数就是空接口实现

	*/
	var a1 B = Cat1{"花猫"}
	var a2 B = Person1{"网二", 26}
	var a3 B = "haha"
	var a4 B = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	test1(a1)
	test1(a2)
	test1(3.15)

	fmt.Println("interface demo2")
}

func test1(a B) {

}

//空接口 任何类型都箱单与实现了空接口
type B interface {
}

type Cat1 struct {
	color string
}

type Person1 struct {
	name string
	age  int
}
