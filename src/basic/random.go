package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	num1 := rand.Int()
	fmt.Println(num1)

	//	没有设置种子之前，每次运行才生的随机数都是固定的
	//	只有设置种子之后随机数才会改变
	for i := 0; i < 10; i++ {
		//只在 10范围内才生随机数
		num := rand.Intn(10)
		fmt.Println(num)
	}
	rand.Seed(1)
	num2 := rand.Intn(10)
	fmt.Println(num2)

	t1 := time.Now()
	timesnap1 := t1.Unix()
	fmt.Println(timesnap1)
	timesnap2 := t1.UnixNano()
	fmt.Println(timesnap2)

	//	为了使种子不停的变化，可以使用时间戳
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		//	step2 调用生成随机数的函数
		fmt.Println("-->", rand.Intn(100))
	}

}
