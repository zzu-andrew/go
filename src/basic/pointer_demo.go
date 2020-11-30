package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("%p\n", &a)

	var p *int
	p = &a
	fmt.Printf("%p\n", p)
	fmt.Println(*p)

	var point **int
	point = &p
	fmt.Println(**point)

	/*
		数组指针： 首先是一个之指针，一个数组的地址
		指针数组： 首先是一个数组， 存储的数据类型是指针
	*/
	// 创建一个普通的数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)

	//创建一个指针， 存储数组的地址--> 数组指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)
	fmt.Println(*p1)
	fmt.Println(&p1)
	// 这里是没有简写， 正常的数组指针调用的方法
	(*p1)[0] = 100
	fmt.Println(arr1)
	p1[0] = 200 // 数组指针的简写形式
	fmt.Println(arr1)

	// 指针数组
	e := 1
	b := 2
	c := 3
	d := 4
	// 普通数组
	arr2 := [4]int{e, b, c, d}
	arr3 := [4]*int{&e, &b, &c, &d}
	fmt.Println(arr2, arr3)
	for i := 0; i < len(arr3); i++ {
		fmt.Println(*arr3[i])
	}
	// 遍历
	for i, v := range arr3 {
		fmt.Println(i, *v)
	}

	fmt.Println("pointer demo")
}
