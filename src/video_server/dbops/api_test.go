package dbops

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func clearTables() {
	_, err := dbConn.Exec("truncate users")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	_, err = dbConn.Exec("truncate video_info")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	_, err = dbConn.Exec("truncate comments")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	_, err = dbConn.Exec("truncate sessions")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUserCredential)
	t.Run("Get", testGetUserCredential)
	t.Run("Del", testDeleteUser)
	t.Run("reget", testRegetUser)
}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("andrew", "123")
	if err != nil {
		t.Error("error of add user")
	}
}

func testGetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("andrew")
	if err != nil {
		t.Error("get user name failed.")
	}
	fmt.Println("password is:", pwd)
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("andrew", "123")
	if err != nil {
		t.Error("del user failed.")
	}
}
func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("andrew")
	if err != nil {
		fmt.Println("Error of RegetUser!")
	}
	if pwd != "" {
		t.Error("delete user test failed.")
	}
}
