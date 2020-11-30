package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql",
		"root:Wang.159@tcp(127.0.0.1:3306)/video_server?charset=utf8")
//	db, err := sql.Open("mysql", "andrew:Wang.159@tcp(127.0.0.1:3306)/test?charset=utf8")

	stmtIns, err := db.Prepare("INSERT INTO users(login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return
	}
	result, err := stmtIns.Exec("andrew", "123")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)


	stmtOut, err := db.Prepare("SELECT pwd FROM users WHERE login_name=?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var pwd string
	err = stmtOut.QueryRow("andrew").Scan(&pwd)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(pwd)
	err = stmtIns.Close()

	if err != nil {
		fmt.Println(err.Error())
	}
	if err != nil {
		fmt.Println("错误信息：", err)
	}
	fmt.Println("连接成功：", db)
	errC := db.Close()
	if err != nil {
		fmt.Println(errC);
	}
	fmt.Println("mysql test")
}
