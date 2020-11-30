package main

import (
	"fmt"
	"reflect"
)



type personStruc struct {
	Name string
	Age  int
	Sex  string
}

func (p personStruc) Say(msg string) {
	fmt.Println("hello:", msg)
}

func (p personStruc) PrintInfo() {
	fmt.Printf("姓名：%s, 年龄：%d, 性别：%s\n", p.Name, p.Age, p.Sex)
}

func main() {

	p1:= personStruc{"ruby", 20, "男"}
	value := reflect.ValueOf(p1)
	fmt.Printf("kind:%s, type: %s\n", value.Kind(), value.Type()) // kind:struct, type: main.personStruc

	// 通过反射对象 并进行函数调用
	methodValue := value.MethodByName("PrintInfo")
	fmt.Printf("kind %s, type %s\n", methodValue.Kind(), methodValue.Type()) // kind func, type func()

	// 没有参数进行调用
	methodValue.Call(nil) // 没有参数直接写Nil或者传入一个空的切片

	args1 := make([]reflect.Value, 0)
	methodValue.Call(args1)

	//有参数
	methodValue2 := value.MethodByName("Say")
	args2:= []reflect.Value{reflect.ValueOf("反射机制")}
	methodValue2.Call(args2)
	fmt.Println("reflect interface")
}
