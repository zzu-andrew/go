package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"video_server/defs"
	"video_server/utils"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users(login_name, pwd) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Add user prepare failed.")
		return err
	}
	result, err := stmtIns.Exec(loginName, pwd)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
	err = stmtIns.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name=?")
	if err != nil {
		fmt.Println("get  user prepare failed", err.Error())
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err.Error())
		return "", err
	}
	err = stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		fmt.Println("delete user failed", err.Error())
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = stmtDel.Close()
	return err
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info(id, author_id, name, display_ctime) VALUES(?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	err = stmtIns.Close()
	if err != nil {
		fmt.Println("close sql db failed.")
	}
	return res, nil
}
