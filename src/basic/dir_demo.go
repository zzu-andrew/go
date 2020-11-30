package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	//
	listFiles(".", 0)
}

func listFiles(dirName string, level int) {
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|  " + s
	}
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range fileInfos {
		fileName := dirName + "/" + file.Name()
		fmt.Printf("%s%s\n", s, fileName)
		if file.IsDir() {
			listFiles(fileName, level+1)
		}
	}
}
