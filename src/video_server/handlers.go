package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	n, err := io.WriteString(w, "create user handler")
	if err != nil {
		fmt.Println(n)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	n, err := io.WriteString(w, "get user name ok")
	if err != nil {
		fmt.Println(n)
	}
	fmt.Println(uname)
}
