package main

import (
	"fmt"
)

func print1to20() int {
	res := 0
	for i := 0; i <= 20; i++ {
		res += i
	}
	return res
}

/*func fistPage(w http.ResponseWriter, r *http.Request) {
	n, err := io.WriteString(w, "<h1>Hello, this is my first page!</h1>")
	if err != nil {
		fmt.Println(n)
	}
}*/

func main() {

	//数组的使用
	// 创建数组
	var arr1 [4]int
	arr1[0] = 1
	arr1[1] = 2
	//数组的长度和容量
	// 因为数组是定长的 因此长度和容量是相同的
	fmt.Println("数组的长度", len(arr1))
	fmt.Println("数组的容量", cap(arr1))
	// 数组的其他创建方式
	var a = [4]int{1, 2, 3, 4}
	fmt.Println(a)

	var b = [5]int{1, 2, 3}
	fmt.Println(b)
	var c = [5]int{1: 1, 3: 2}
	fmt.Println(c)
	//自动推断数组的长度
	d := [...]int{1, 2, 3, 4}
	fmt.Println(d)
	var f = [...]int{1, 2, 34, 5}
	fmt.Println(f)

	//-------------------------------------
	/*http.HandleFunc("/", fistPage)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("err.")
	}*/

	//fmt.Println(runtime.NumCPU())
	fmt.Println("array test")
}
