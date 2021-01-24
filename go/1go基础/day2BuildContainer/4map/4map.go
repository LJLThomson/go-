package main

import "fmt"

/**
1、map定义初始化，map时Hashmap,无序的
2、map遍历
3、key拼错，会输出什么了
4、删除元素
*/
func main() {
	//1定义,map[key]v, map[key]map[key][v]
	m := map[string]string{
		"name":    "day2",
		"curse":   "golang",
		"site":    "baidu",
		"quality": "good"}
	m2 := make(map[string]int) //empty map
	var m3 map[string]int      //nil,在go中也可以安全使用
	fmt.Println(m)
	fmt.Println(m2)
	fmt.Println(m3)
	//2遍历,map是HashMap，无序的
	for k, v := range m {
		fmt.Println(k, v)
	}
	//3拼错会出现什么，空串
	curseName := m["curse"]
	fmt.Println(curseName)
	curseName2, ok := m["cousxxxx"]
	fmt.Println(curseName2, ok)
	if curseName3, ok := m["name"]; ok {
		fmt.Println(curseName3)
	} else {
		fmt.Println("key not exist")
	}
	//4 删除元素
	delete(m, "name")
	if curseName2, ok = m["name"]; ok {
		fmt.Println(curseName2)
	} else {
		fmt.Println("key not exist")
	}
}
