package main

import (
	"fmt"
	mytest "l_package/PK2"
	"l_package/person"
	"l_package/pk1"
	"l_package/utils"
	"l_package/utils/timeutils"
)

//import "./utils"

//默认从gopath goroot下的src下开始找 -- 绝对路径
//import "l_package/utils"
//相对路径
//import "./utils"

// 如果导入包只是为了执行包里面的init函数， 导入的时候需要在包名前面加上 _ 表示导入但是不使用
// 加上  . 的意思是包中的函数不需要加前缀直接在当前环境中使用
func main() {
	/*
		关于包的使用：
		1. 一个目录下的统计文件归属一个包，package的声明要一致
		2.包的名字和目录的名字可以不同，但是通常保持一致
		3. 包可以嵌套
		4. 同包下的函数不需要导入可以直接使用
		5. main包特殊，其他包不能导入
		6.路径要从src下开始写，有几层就写几层
		7. 一个包可以被导入多此，但是初始化只一次
	*/
	utils.Count()
	timeutils.PrintPackage()
	pk1.MyTest()
	mytest.MyTest2()
	fmt.Println("===================")
	p1 := person.Person{Name: "xiaming", Sex: "男"}
	fmt.Print(p1)

}
