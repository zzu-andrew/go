package main

import (
	"bufio"
	"fmt"
	"os"
)

type animal struct {
	eye  int
	foot int
}

type dog struct {
	animal
	watch string
}

func (a *animal) animalInit() error {
	return nil
}

func (d *dog) Watch() error {
	fmt.Printf("%s\n", d.watch)
	return nil
}

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

	var aDog dog
	aDog.eye = 2
	aDog.watch = "wang wang"
	aDog.foot = 4
	_ = aDog.Watch()

}
