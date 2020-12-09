package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		time
	*/
	t1 := time.Now()
	fmt.Println(t1)
	fmt.Println(t1.Hour())

	// 格式化字符使用 2006 01 02 15：04：05
	// 如果使用的不是  就会按照原字符输出
	s1 := t1.Format("2006/01/02")
	fmt.Println(s1)
	s2 := t1.Format("2006  01  02  15 ssss 08")
	fmt.Println(s2)
}
