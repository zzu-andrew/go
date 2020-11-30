package main

import (
	"fmt"
	"math"
)

func main() {
	var t1 Triangle = Triangle{3, 4, 5}
	fmt.Println(t1.peri())
	fmt.Println(t1.area())
	fmt.Println(t1.a, t1.b, t1.c)

	var c1 Circle = Circle{4}
	fmt.Println(c1.peri())
	fmt.Println(c1.area())
	fmt.Println(c1.redius)

	// 接口实现调用， 实现对数据的分装
	var s1 Shape
	s1 = t1
	fmt.Println(s1.peri())
	fmt.Println(s1.area())

	var s2 Shape
	s2 = c1
	fmt.Println(s2.peri())
	fmt.Println(s2.area())

	testShape(t1)
	testShape(c1)
	testShape(s1)

	getType(t1)
	getType(c1)
	getType(s1)

	// 指针是指针的类型
	var t2 *Triangle = &Triangle{3, 4, 5}
	fmt.Printf("t:%T, %p\n", t2, &t2)
	getType(t2)
	getType2(t2)

	fmt.Println("interface demo 4")
}

func getType2(s Shape) {
	switch ins := s.(type) {
	case Triangle:
		fmt.Println("三角形", ins.a, ins.b, ins.c)
	case Circle:
		fmt.Println("圆形", ins.redius, ins.area(), ins.peri())
	case *Triangle:
		fmt.Println(&ins)
	}
}

//定义一个接口
type Shape interface {
	peri() float64
	area() float64
}

// 定义实现类 ：三角形
type Triangle struct {
	a, b, c float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := t.peri() / 2
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

type Circle struct {
	redius float64
}

func (c Circle) peri() float64 {
	return c.redius * 2 * math.Pi
}

func (c Circle) area() float64 {
	return math.Pow(c.redius, 2) * math.Pi
}

func testShape(s Shape) {
	fmt.Printf("周长：%.2f, 面积：%.2f\n", s.peri(), s.area())
}

func getType(s Shape) {
	if ins, ok := s.(Triangle); ok {
		fmt.Println("是三角形，三边是", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println(ins.redius, ins.peri(), ins.area())
	} else if ins, ok := s.(*Triangle); ok { // 这里是重新生成了数据，只是和传入的数据一样
		fmt.Printf("ins; %T, %p\n", ins, &ins)
	} else {
		fmt.Println("我也不知道了")
	}
}
