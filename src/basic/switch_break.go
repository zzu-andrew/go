package main

import "fmt"

func main() {
	/*
		 switch 中的break和 fallthrough 语句
		switch 语句中默认每个case后面会加上break
		break可以在switch中也可以在for循环中
	*/
	n := 1
	switch n {
	case 1:
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
		break
		fmt.Println("这里将不会被打印")
	case 2:
		fmt.Println("4")
	default:
		fmt.Println("数据出错")
	}

	//fallthrough 代表者穿透的意思，上一个case匹配成功将紧接着执行fallthrough后面的case
	switch n {
	case 1:
		fmt.Println("case 1 匹配成功")
		fallthrough
	case 2:
		fmt.Println("fallthrough 后面的才能进入的case")
	case 3:
		fmt.Println("没有 fallthrough 的case ")
	default:
		fmt.Println("default")
	}

}
