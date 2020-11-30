package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile := "test.txt"
	desFile := "des.txt"
	file1, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T", file1)
	n1, err := file1.WriteString("string test")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(n1)
	n, err := copyFIle1(srcFile, desFile)
	if err != nil {
		fmt.Println("文件复制出错")
	}
	fmt.Println(n)
	fmt.Println("io demo copy file")
}

func copyFIle1(srcFile, desFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(desFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	bs := make([]byte, 1024, 1024)
	n := -1
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("copy end")
			break
		} else if err != nil {
			fmt.Println("copy failed")
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil
}
