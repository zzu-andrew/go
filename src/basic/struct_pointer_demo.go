package main

import "fmt"

func main() {

	// 使用new函数， 是go语言中专门用于创建某种指针的函数
	p1 := personStruct{
		age:     27,
		name:    "xiaoming",
		address: "杭州市",
		sex:     "男",
	}
	fmt.Println(p1)
	p2 := new(personStruct) // &{0   } 是个结构体类型的指针
	fmt.Println(p2)
	// 指针在C语言中必须使用 * 或者  -> 进行解引用，在Go语言中不需要
	//可以直接用这种形式调用指针中的数据
	p2.sex = "男"
	p2.name = "小明"
	p2.age = 27
	p2.address = "杭州市"
	fmt.Println(p2) //&{27 小明 杭州市 男} 指针类型，并打印出指针的内容

	// 结构体中的匿名字段
	st1 := struct {
		name string
		age  int
	}{"xiaoming", 26}
	fmt.Println(st1)
	// 匿名字段结构体的使用
	w1 := Worker{"李晓华", 32}
	// 没有字段名，就默认将字符类型作为类型名字
	// 其实内部编译器实现也是按照  string : string实现，
	// 因此匿名结构体中的类型不能重复
	fmt.Println(w1.string)

	fmt.Println("struct point demo")
}

type personStruct struct {
	age     int
	name    string
	address string
	sex     string
}

// 带有匿名字段的结构体
type Worker struct {
	string
	int
}
