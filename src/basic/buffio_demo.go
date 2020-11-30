package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "test.jpeg"
	fmt.Println(fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//	创建reader对象
	b1 := bufio.NewReader(file)
	p := make([]byte, 1024)
	n1, err := b1.Read(p)
	fmt.Println(n1)
	fmt.Println(string(p[:n1]))


}
