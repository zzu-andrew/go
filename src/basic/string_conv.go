package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		strconv包 自读出啊基本类型之间的转换
	*/
	s1 := "true"
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T, %t\n", b1, b1)

	s2 := "100"
	// bas 代表的是按照几进制进行转换
	i2, err := strconv.ParseInt(s2, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T,%d\n", i2, i2)

	i3, _ := strconv.Atoi("-42")
	fmt.Printf("%T, %d\n", i3, i3)

	s3 := strconv.Itoa(-46)
	fmt.Printf("%T, %s\n", s3, s3)
}
