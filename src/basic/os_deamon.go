package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	fileInfo, err := os.Stat("/home/andrew/go/src/basic/test.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("%T\n", fileInfo)

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	// 文件路径
	fileName := "/home/andrew/go/src/basic/test.txt"
	fileName2 := "test.txt"
	fmt.Println(filepath.IsAbs(fileName))
	fmt.Println(filepath.IsAbs(fileName2))
	fmt.Println(filepath.Abs(fileName2))

	// 要想在当前目录下创建文件，编译配置也需要更改到当前目录下
	file1, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file1.Name())
	n, err := file1.WriteString("test file")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	b := make([]byte, 9, 100)
	// 指定开始的位置，否则读出来的文件是空的，因为写完之后光标在文件结尾
	n2, err := file1.ReadAt(b, 0)
	fmt.Println(n2)
	fmt.Println(b)

	fmt.Println("os demon")
}
