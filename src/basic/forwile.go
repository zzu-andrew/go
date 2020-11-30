package main

import (
	"fmt"
	"math"
)

func main() {

	//	for init; condition; post{}
	var i int = 1
	for n := i; n <= 5; n++ {
		fmt.Println("hello world")
	}
	//	省略初始化和post
	// 其实C语言中也支持，只是不支持省略分号但是go中分号也能省略
	j := 1
	for j <= 5 {
		fmt.Println("j = ", j)
		j++
	}
	// 省略所有的for参数，相当于 while(1)
	k := 1
	for {
		k++
		fmt.Println(k)
		if k > 100 {
			break
		}
	}
	k = 1
	for {
		k++
		if k == 9 {
			continue
		}
		fmt.Println(k)
		if k > 10 {
			break
		}
	}
	//	水仙花数 水仙花数，就是个位 十位 百位 三次幂的和本身大小一样，就是水仙花数
	for k = 100; k < 1000; k++ {
		x := k / 100
		y := k / 10 % 10
		z := k % 10
		w := math.Pow(float64(x), 3) + math.Pow(float64(y), 3) + math.Pow(float64(z), 3)
		if int(w) == k {
			fmt.Println(k)
		}
	}

}
