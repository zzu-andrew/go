package main

import "fmt"

func main() {
	/*
		// 方法调用，除了对应类型的数据，也可以是对应类型的指针
		// 方法 需要执行调用者类型的特殊函数 一个方法就是包含接受者的函数
		方法名是可以冲突的，只要调用名不一致就行
		方法：是某一个类别的行为功能，需要制定的接受者调用
		可以方法名相同，只要接受者不同就可以了
		func (t type) methodName(parameter list)(return list){

		}
		func funcname(parameter list)(return list){

		}*/
	w1 := WorkerStr{name: "xiaoming", age: 30, sex: "男"}
	w1.work()
	p1 := &w1
	fmt.Println(p1)
	p2 := &WorkerStr{name: "李四", age: 26, sex: "男"}
	fmt.Printf("%T\n", p2) // *main.WorkerStr 指针类型
	p2.work()

	p2.rest()
	p2.printInfo()

	c1 := Cat{age: 26, color: "white"}
	c1.printInfo()

	fmt.Println("method demo")

}

type WorkerStr struct {
	//	字段
	name string
	age  int
	sex  string
}

type Cat struct {
	color string
	age   int
}

// 定义方法, 调用者可以在方法中直接使用
func (w WorkerStr) work() {
	fmt.Println(w.name, "正在工作")
}

func (w *WorkerStr) rest() {
	fmt.Println(w.name, "rest.")
}

func (p *WorkerStr) printInfo() {
	fmt.Println(p.name, p.sex, p.age)
}
func (p *Cat) printInfo() {
	fmt.Println(p.color, p.age)
}
