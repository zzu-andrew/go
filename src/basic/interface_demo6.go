package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	_ "path/filepath"
	"reflect"
	"strconv"
	"strings"
)

//conf.ini
/**********************************************
; mysql config
[mysql]
address=10.17.112.22
port=3306
username=root
password=root


[redis]
host=127.0.0.1
port=6379
password=root
database=0

***********************************************/

type mySQL struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

// 对应配置表中的字段名
type Config struct {
	mySQL       `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(filename string, data interface{}) (err error) {

	idx := 0
	// 取指针  reflect.TypeOf(data)
	// 若是想取数值需要使用  reflect.ValueOf()
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	// 传进来的类型必须是指针类型
	if t.Kind() != reflect.Ptr {
		err = fmt.Errorf("data may be err")
		//err  = errors.New("data should be ptr ")
		return
	}
	// data必须是结构体指针
	if t.Elem().Kind() != reflect.Struct {
		err = fmt.Errorf("data may be not struct")
		//err  = errors.New("data should be ptr ")
		return
	}
	// 读文件 得到字节类型数据

	f, err := os.Open("conf.ini")
	if err != nil {
		fmt.Println(err.Error())
	}
	var structName string
	//建立缓冲区，把文件内容放到缓冲区中
	buf := bufio.NewReader(f)
	var valueBreak int
	for valueBreak == 0 {
		//遇到\n结束读取
		b, errR := buf.ReadBytes('\n')
		if errR != nil {
			line := string(b)
			fmt.Println("=========================================================", line)
			if errR == io.EOF {
				// 当读取到最后一行，把退出的标记进行记录
				valueBreak++
			}
			fmt.Println(errR.Error())
		}
		//fmt.Println(string(b))
		// 去除空格
		line := string(b)
		idx++
		// 空行跳过
		if len(line) == 0 {
			continue
		}
		// 单个换行符的行业跳过
		if line[0] == '\n' {
			continue
		}
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			//continue
		}
		fmt.Println(line)
		// 如果 [开头 就是节
		if strings.HasPrefix(line, "[") {
			//fmt.Println("line =========", line)
			if line[0] != '[' && line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}

			sectionName := strings.TrimSpace(line[1 : len(line)-2])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d sybtax error", idx+1)
				return
			}
			//	 根据字符串，去data里面根据反射找到对应的结构体

			for i := 0; i < t.Elem().NumField(); i++ {

				filed := t.Elem().Field(i)
				fmt.Println("=====================================================", sectionName)
				if sectionName == filed.Tag.Get("ini") {

					structName = filed.Name
					fmt.Println(sectionName+" 对应的结构体是：", structName)

				}
			}
			fmt.Println("_____________________________________________________")

		} else {
			// 如果不是 [ 开头就是键值对
			//fmt.Println("not space line == ", line)
			//fmt.Printf("%d\n", []byte(line))
			// 1. 以等号分割这一行，等号左边是Key, 等号右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line %d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			fmt.Println("get index = ", index)
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//fmt.Println("key =", key, "value  = ", value)
			//	2. 根据struct name去data中把对应的嵌套结构体取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 拿到嵌套结构体的值信息
			sType := sValue.Type()                     // 拿到嵌套结构体的类型信息
			fmt.Println(sValue, sType)
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("structObj 应该是一个结构体")
				return
			}
			//	3. 遍历嵌套结构体中的每一个字段，判断tag是不是等于key
			var fieldName string
			var fileType reflect.StructField
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i)
				fileType = field
				if field.Tag.Get("ini") == key {
					//	 找到了对应的字段
					fieldName = field.Name
					break
				}
			}
			//	4.如果key == tag,给这个字段赋值
			// 根据filedName 去取出这个字段进行赋值
			fileObj := sValue.FieldByName(fieldName)
			//fmt.Println("+++++++++++++++++++++++++++++++++++++++", fileObj)
			//fmt.Println(fieldName)
			// 对其进行赋值
			fmt.Println(fieldName, fileType.Type.Kind())
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				valueInt, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				valueBool, err := strconv.ParseBool(value)
				if err != nil {
					return nil
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				valueFloat, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return nil
				}
				fileObj.SetFloat(valueFloat)

			}

		}

	}

	/*content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	//_ = ioutil.WriteFile("test.ini", content, 0644)
	//将文件内容转换成字符串
	lineStr := strings.Split(string(content), " ")
	//fmt.Println(lineStr)
	// 一行一行读数据*/
	/*for idx, line := range lineStr {

	// 去除空格
	line = strings.TrimSpace(line)
	fmt.Println(line, "idx = ", idx)
	if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
		//continue
	}
	//fmt.Println(line)
	// 如果 [开头 就是节
	if strings.HasPrefix(line, "[") {
		fmt.Println("line =========",line)
		if line[0] != '[' && line[len(line)-1] != ']' {
			err = fmt.Errorf("line:%d syntax error", idx+1)
			return
		}
		sectionName := strings.TrimSpace(line[1 : len(line)-1])
		if len(sectionName) == 0 {
			err = fmt.Errorf("line:%d sybtax error", idx+1)
			return
		}
		//	 根据字符串，去data里面根据反射找到对应的结构体
		v := reflect.ValueOf(data)
		for i := 0; i < t.Elem().NumField(); i++ {
			filed := t.Elem().Field(i)
			if sectionName == filed.Tag.Get("ini") {
				structName := filed.Name
				fmt.Println("test", structName)

			}
		}

		fmt.Println(v)

	} else {
		// 如果不是 [ 开头就是键值对
		//fmt.Println("==========")
	}*/

	return nil
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Println("interface demo reflect")
	}

	fmt.Println(cfg)
}
