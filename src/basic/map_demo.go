package main

import "fmt"

func main() {
	// map也是引用类型，和切片一样长度可变

	//创建 map
	var map1 map[int]string         // 只有声明没有初始化  是一个nil的map
	var map2 = make(map[int]string) // 这个map非空 但是没有放数组
	var map3 = map[string]int{
		"GO":     97,
		"Python": 89,
		"Java":   79,
		"Html":   78,
	}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)
	if map1 == nil {
		map1 = make(map[int]string)
	}

	// 存储键值对到map中去
	map1[0] = "wangxiaomei"
	map1[1] = "daxiong"
	map1[2] = "wangdada"
	map1[3] = "lisi"
	fmt.Println(map1)

	// 根据key只获取对应的value
	fmt.Println(map1[2])
	fmt.Println(map1[30])  // 不存在的只，或取出来是value类型的默认值

	// value key
	for index, value := range map1{
		fmt.Println(index, value)
	}

	v1:= map1[2]
	fmt.Println(v1)

	// 删除数据根据key来删除，如果key存在就删除， 如果不存在就不删除
	delete(map1, 2)
	fmt.Println(map1)
	delete(map1, 30)
	fmt.Println(map1)

	fmt.Println("map demo")
}
