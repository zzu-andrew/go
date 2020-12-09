package main

import (
	"fmt"
	"reflect"
)

type personStr struct {
	Name string
	Age  int
	Sex  string
}

func (p personStr) Say(msg string) {
	fmt.Println("hello:", msg)
}

func (p personStr) PrintInfo() {
	fmt.Printf("姓名：%s, 年龄：%d, 性别：%s\n", p.Name, p.Age, p.Sex)
}

func main() {
	/*
		反射 通过发射可以获取一个接口类型变量的类型和数量
	*/

	var x float64 = 3.14
	fmt.Println("type:", reflect.TypeOf(x))        // type: float64
	fmt.Println("value of x:", reflect.ValueOf(x)) // value of x: 3.14

	fmt.Println("================================")
	// 根据反射出来的值，来获取对应的类型和取值
	v := reflect.ValueOf(x)
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("type:", v.Type())
	fmt.Println("value:", v.Float())
	fmt.Println("-------------------------------")
	//	 已知类型的获取
	var num float64 = 1.23
	value := reflect.ValueOf(num)
	//	已知类型可以直接转化
	convertValue := value.Interface().(float64)
	fmt.Println("convertValue:", convertValue)
	fmt.Println("-------------------------------------")
	p1 := personStr{Name: "xiaohua", Age: 30, Sex: "男"}
	GetMessage(p1)
}
func GetMessage(input interface{}) {
	getType := reflect.TypeOf(input) // 先获取类型
	fmt.Println("get Type is ", getType.Name())
	fmt.Println("get kind of :", getType.Kind())

	getValue := reflect.ValueOf(input)
	fmt.Println("get value is :", getValue)
	/*
		step1 先获取type对象： reflect.Type
		NumFiled
		Field(index)
		step2 通过filed获取每一个filed字段
		step3 Interface() 得到对应value

	*/
	for i := 0; i < getType.NumField(); i++ {
		filed := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("字段名：%s, 字段类型：%s, 字段数据： %v \n", filed.Name, filed.Type, value)
	}

	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名称：%s, 方法类型：%v", method.Name, method.Type)
	}

}
