package main

import "fmt"

func main() {
	/*
		可变参数:
		func myFunc(arg ...int){}
	*/

	getSum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	// 切片
	// 如果有可变参数， 可变参数要放到函数参数的最后
	// 一个函数的参数列表中，最多只能有一个可变参数
	// 要是传入的是切片需要将切片后面加上三个...， 这样函数才能知道传入的是切片
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	getSum(s1...)
}

func getSum(nums ...int) {
	// 可以看到，可变参数其实是一个切片
	// 调用函数的时候，可以传入一个或者多个参数
	// print类型函数， 空接口实现的任意可变参数
	fmt.Printf("nums = %T\n", nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println(sum)
}
