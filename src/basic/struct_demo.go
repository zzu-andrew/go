package main

import "fmt"

func main() {

	var p1 person
	p1.address = "杭州"
	p1.age = 27
	p1.name = "xiaoming"
	p1.sex = "男"
	fmt.Println(p1)
	p2 := person{}
	fmt.Println(p2)

	fmt.Println("struct demo")
}

//Person
type person struct {
	name    string
	age     int
	sex     string
	address string
}
