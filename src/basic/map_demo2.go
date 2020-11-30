package main

import (
	"fmt"
	"sort"
)

func main() {
	// 使用for range进行遍历
	map1 := make(map[int]string)
	map1[1] = "GO"
	map1[2] = "Python"
	map1[3] = "Java"
	map1[4] = "C"
	map1[5] = "C++"
	fmt.Println(map1)

	//遍历并不是按照顺序进行的
	// 也就是说map是无序的
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	fmt.Println("test map1")
	// 按顺序只能自己指定
	for i := 1; i <= len(map1); i++ {
		fmt.Println(i+1, map1[i])
	}
	// 获取所有的key
	// 进行排序
	// 进行遍历
	keys := make([]int, 0, len(map1))
	fmt.Println(keys)
	// 获取素有的key值
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	sort.Ints(keys)
	fmt.Println(keys)
	for _, key := range keys {
		fmt.Println(key, map1[key])
	}
	//
	s1 := []string{"Apple", "Windows", "Orange", "abc", "王", "acd", "acc"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)
	fmt.Println("map demo")
}
