package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Rmb   float64
	Hobby []string
	Sex   bool
}

/**
1、student序列化
2、student对象转化为json序列化
*/
func main1() {
	student := Student{"张三", 25, 12345.67, []string{"抽烟", "喝酒", "赌博"}, true}
	bytes, err := json.Marshal(student)
	if err != nil {
		fmt.Println("序列化转化错误")
	} else {
		fmt.Println(string(bytes))
	}
}

/**
map转化为json
*/
func main2() {
	dateMap := make(map[string]interface{})
	dateMap["name"] = "张三"
	dateMap["age"] = 25
	dateMap["rmb"] = 123.45
	dateMap["hobby"] = []string{"抽烟", "喝酒", "赌博"}
	dateMap["sex"] = true
	bytes, err := json.Marshal(dateMap)
	if err != nil {
		fmt.Println("序列化转化错误")
	} else {
		fmt.Println(string(bytes))
	}
}

/**
slice切片转json
*/
func main() {
	dateMap1 := make(map[string]interface{})
	dateMap1["name"] = "张三"
	dateMap1["age"] = 25
	dateMap1["rmb"] = 123.45
	dateMap1["hobby"] = []string{"抽烟", "喝酒", "赌博"}

	dateMap2 := make(map[string]interface{})
	dateMap2["name"] = "张三"
	dateMap2["age"] = 25
	dateMap2["rmb"] = 123.45
	dateMap2["hobby"] = []string{"抽烟", "喝酒", "赌博"}

	dateMap3 := make(map[string]interface{})
	dateMap3["name"] = "张三"
	dateMap3["age"] = 25
	dateMap3["rmb"] = 123.45
	dateMap3["hobby"] = []string{"抽烟", "喝酒", "赌博"}

	dateSlice := make([]map[string]interface{}, 0)
	dateSlice = append(dateSlice, dateMap1, dateMap2, dateMap3)
	bytes, err := json.Marshal(dateSlice)
	if err != nil {
		fmt.Println("序列化转化错误")
	} else {
		fmt.Println(string(bytes))
	}
}
