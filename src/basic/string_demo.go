package main

import (
	"fmt"
	"strings"
)

func main() {

	var s1 string = "name test"
	fmt.Println(s1)
	fmt.Println(len(s1))
	var s2 string = "整个南国"
	fmt.Println(s2)
	fmt.Println(len(s2))

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c\n", s1[i])
	}

	for i, v := range s1 {
		fmt.Println(i, v)
	}

	fmt.Println(strings.Contains(s1,"s"))


	fmt.Println("string demo")
}
