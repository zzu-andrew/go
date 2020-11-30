package main

import "fmt"

func main() {
	/*
		 每一个切片引用了一个底层数组
		切片本身不存储任何数据，都是底层数据结构存储，所以修改切片就是修改这个数组中的数据
		当向切片中添加数据时，如果没有超过容量，直接添加，若果超过容量自动扩容(成倍增长)
		切片一旦扩容个，就会指向一个新的底层数组
		 *
	*/
	// 底层创建一个数组， 切片s1引用数组
	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)
	s1 = append(s1, 4, 5)
	// 成倍增长之后 长度应该为 5 容量会变成6
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 4, 5, 4, 7, 8, 9)
	// 成倍增长之后 长度应该为 5 容量会变成6
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 4, 5, 6)
	// 一直按照成倍的增长，当长度为12时，在增长就是24了
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1)

	// 调试打印
	fmt.Println()
	fmt.Println("slice demo2")
}
