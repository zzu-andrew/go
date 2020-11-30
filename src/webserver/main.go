package main

import (
	"fmt"
	"io"
	"net/http"
)

func print1to20() int {
	res := 0
	for i := 0; i <= 20; i++ {
		res += i
	}
	return res
}

func fistPage(w http.ResponseWriter, r *http.Request) {
	n, err := io.WriteString(w, "<h1>Hello, this is my first page!</h1>")
	if err != nil {
		fmt.Println(n)
	}
}

func main() {
	//-------------------------------------
	http.HandleFunc("/", fistPage)
	err := http.ListenAndServe(":8000", nil)
	if err != nil{
		fmt.Println("err.")
	}
}
