package main

import (
	"fmt"
)

func main() {

	// go语言中switch的特殊用法
	// 当没有expr的时候，默认是bool类型，并直接定位到true上
	// 即省略 switch后的变量，相当于直接作用在true上
	switch {
	case true:
		fmt.Println("is true")
	case false:
		fmt.Println("si false")
	}

	//	使用省略变量之后只到true上的特性可以实现
	/*
		成绩：
		[0-59]，不及格
		[60-79]，及格
		[80-89]，中
		[90-100]，优秀
	*/
	score := 89
	switch {
	case score >= 0 && score < 60:
		fmt.Println("不及格")
	case score >= 60 && score < 70:
		fmt.Println("及格")
	case score >= 70 && score < 80:
		fmt.Println("中")
	case score >= 80 && score < 90:
		fmt.Println("良")
	case score >= 90 && score < 100:
		fmt.Println("优秀")
	default:
		fmt.Println("成绩有误")
	}

	//	支持多种case放到一行上
	letter := ""
	switch letter {
	case "A", "B", "C":
		fmt.Println("case 1")
	case "M", "N":
		fmt.Println("case2")
	default:
		fmt.Println("other")
	}

	//	带有初始化语句
	switch language := "golang"; language {
	case "golang":
		fmt.Println("是go语言")
	case "java":
		fmt.Println("java语言")
	case "python":
		fmt.Println("python语言")
	}

}
