package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn  *sql.DB
	err    error
)

func init() {
	db, err := sql.Open("mysql",
		"root:Wang.159@tcp(127.0.0.1:3306)/video_server?charset=utf8")
	if err != nil {
		// 无法正常进行业务逻辑
		fmt.Println(err.Error())
		panic(err.Error())
	}
	//fmt.Printf("%T", dbConn)
	dbConn = db
	fmt.Println("init mysql")
}
