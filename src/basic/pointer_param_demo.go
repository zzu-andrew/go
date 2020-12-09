package main

import "fmt"

func main() {
	a := 100
	fun2(&a)
	fmt.Println(a)

	fmt.Println("pointer param")
}

func fun2(p1 *int) {
	fmt.Println(*p1)
	*p1 = 30
}
