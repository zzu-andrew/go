package main

import "fmt"

func main() {

	// 切片中的数据 [start, end)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)
	// 切片中的数据 [start, end)
	// 因为是开区间，所以取出来的下标是4结束， 0开始4下标刚好存储的是5
	// 开始是从下标开始闭区间，结束时开区间
	slice1 := arr[0:5]
	fmt.Println(slice1)

	slice2 := arr[1:5]
	fmt.Println(slice2)

	slice3 := arr[3:8]
	fmt.Println(slice3)

	slice4 := arr[:]
	fmt.Println(slice4)

	// 存储的地址
	fmt.Printf("%p\n", &arr)
	// 因为切片本身就是一个引用的地址，因此可以直接打印地址
	fmt.Printf("%p\n", slice1)

	// 切片的cap 是从起始位置到结束位置
	fmt.Printf("len = %d,cap = %d\n", len(slice1), cap(slice1))
	fmt.Printf("len = %d,cap = %d\n", len(slice2), cap(slice2))
	fmt.Printf("len = %d,cap = %d\n", len(slice3), cap(slice3))
	// 更改数组的数据
	arr[4] = 100
	fmt.Println(slice1)
	// 更改切片中的数据
	slice1[4] = 5
	fmt.Println(arr)
	// 切片的添加也是一样的效果
	slice1 = append(slice1, 1, 1, 1, 1)
	fmt.Println(arr)
	fmt.Printf("slice1 = %p, arr = %p\n", slice1, &arr)
	// 切片的添加也是一样的效果
	fmt.Println("=======================================")
	slice1 = append(slice1, 1, 1, 1, 1, 1, 1, 1)
	fmt.Println(arr)    // 数组不变
	fmt.Println(slice1) // slice1新申请了块内存
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Printf("slice1 = %p, arr = %p\n", slice1, &arr)

	fmt.Println("slice demo")
}
