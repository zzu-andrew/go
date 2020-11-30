package main


/*
1. 指定使用的数据库，use table_name
2. 显示数据库中的所有表单数据 show tables
3. 解析对应的表单数据 describle comments
mysql> show tables;
+------------------------+
| Tables_in_video_server |
+------------------------+
| comments               |
| sessions               |
| users                  |
| video_info             |
+------------------------+
4 rows in set (0.00 sec)

mysql> describe comments;
+-----------+--------------+------+-----+---------+-------+
| Field     | Type         | Null | Key | Default | Extra |
+-----------+--------------+------+-----+---------+-------+
| id        | varchar(64)  | NO   | PRI | NULL    |       |
| video_id  | varchar(64)  | YES  |     | NULL    |       |
| author_id | int unsigned | YES  |     | NULL    |       |
| content   | text         | YES  |     | NULL    |       |
| time      | datetime     | YES  |     | NULL    |       |
+-----------+--------------+------+-----+---------+-------+
5 rows in set (0.00 sec)
*/
import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)

	return router
}
func main() {
	r := RegisterHandlers()
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println(err)
	}
}
