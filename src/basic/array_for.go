package main

import (
	"fmt"
)

func main() {

	var arr1 [5]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for index, value := range arr1 {
		fmt.Printf("index = [%d], value = [%d]\n", index, value)
	}

	for _, value := range arr1 {
		fmt.Printf("value = [%d]\n", value)
	}
}
