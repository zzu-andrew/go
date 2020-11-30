package main

import "fmt"

func main() {

	//	1. 数组
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	//	切片
	var s1 []int
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4} // 变长
	fmt.Println(s2)
	fmt.Printf("%T,%T\n", arr, s2)
	//	func make(Type, size IntegerType) Type
	//	Type 类型 slice mao chan
	//	size 长度 实际元素存储的数量
	// 容量cap 最多能欧存储的元素个数
	s3 := make([]int, 5, 8)
	fmt.Printf("容量：%d,长度：%d\n", cap(s3), len(s3))
	s3[0] = 1
	s3[1] = 2
	s3[2] = 3
	fmt.Println(s3)

	//	append 的两种用法
	s4 := make([]int, 0, 5)
	fmt.Printf("%p\n", s4)
	s4 = append(s4, 1, 2)
	fmt.Printf("%p\n", s4)
	s4 = append(s4, 3, 4, 5, 6, 7)  // 一旦扩容 长度会城北的增长
	fmt.Printf("容量：%d, 长度：%d\n", cap(s4), len(s4))
	fmt.Printf("%p\n", s4)
	//	添加切片
	s4 = append(s4, s3...)
	fmt.Println(s4)
	fmt.Printf("%p\n", s4)

}
