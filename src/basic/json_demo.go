package main

import (
	"encoding/json"
	"fmt"
)

// json  --  struct
// 注意因为这个结构体要在 @json.Marshal 内部使用，所以想要出来完整的json就必须使用大小的变量
// 当有要求，生成的json必须是小写的时候，可以使用关键字  json指定字段的名字 如`json:"name" db:"name"`
type jsonPerson struct {
	Name string `json:"name" db:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := jsonPerson{
		Name: "xiaohong",
		Age:  23,
	}

	byteStr, _ := json.Marshal(p1)
	fmt.Println(string(byteStr))
	fmt.Printf("%#v", string(byteStr))
}
