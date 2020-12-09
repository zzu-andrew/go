package main

import "fmt"

func main() {
	// 面向对象OOP
	// 1. 创建父类对象
	p1 := Person{name: "zhangsan", age: 30}
	fmt.Println(p1.age, p1.name)
	fmt.Println(p1)
	// 创建子类对象
	s1 := Student{Person{"lisi", 27}, "zheda"}
	fmt.Println(s1)
	// 字段提升，省略中间的匿名字段
	// s1.Person.name ---> s1.name
	// 这就是提升字段， 模拟继承 提升之后，子类中能直接访问父类中的字段
	fmt.Println(s1.name, s1.age, s1.school)

	fmt.Println("oop demo")
}

// 定义父类
type Person struct {
	name string
	age  int
}

// 定义子类
type Student struct {
	Person        // 模拟继承结构
	school string // 子类的新增属性
}
