package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := "conf.ini"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//	创建reader对象
	fileReader := bufio.NewReader(file)
	// 这里可以加个for循环，知道err == io.EOF跳出循环
	fileBuff, isPrefix, err := fileReader.ReadLine()
	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}
	fmt.Println(string(fileBuff), isPrefix, err)

	/*p := make([]byte, 1024)

	n1, err := b1.Read(p)
	fmt.Println(n1)
	fmt.Println(string(p[:n1]))*/


}