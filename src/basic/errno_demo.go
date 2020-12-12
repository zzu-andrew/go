package main

import (
	"errors"
	"fmt"
)

func main() {

	/*
		错误：在可能出现问题的地方出现问题， 如打开一个文件打开失败是意料之中的事情
		异常：在不该出现错误的地方出现错误， 如出现空指针，意料之外的情况
	*/
	//exec.Command("pwd")
	//exec.Command("ps")
	//f, err := os.Open("test.txt")
	//if err != nil {
	//	log.Fatal(err)  // print + exit()
	//	//fmt.Println(err)
	//}
	//fmt.Println(f.Name(), "打开成功.....")
	////-------------------
	/*
		error 内置了数据类型，内置的接口
			定义了方法:Error() string
		// The error built-in interface type is the conventional interface for
		// representing an error condition, with the nil value representing no error.
		type error interface {
			Error() string
		}
	*/
	err1 := errors.New("创建测试的错误处理函数")
	fmt.Println(err1)
	fmt.Printf("%T\n", err1)

	//另外一个创建错误的方法
	err2 := fmt.Errorf("错误的信息：%d", 100)
	fmt.Println(err2)
	fmt.Printf("%T\n", err2)

	fmt.Println("--------------------")
	err3 := checkAge(-30)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	fmt.Println("程序 go on")

	fmt.Println("errno demo")
}
func checkAge(age int) (err error) {
	if age < 0 {
		//return errors.New("年龄不合法")
		err = fmt.Errorf("你给的年龄是%d,不合法", age)
		return
	}
	fmt.Println("年龄是：", age)
	return nil
}
