package main

import (
	"encoding/json"
	"fmt"
)

/**
json反序列化
json转map
*/
func main1() {
	jsonStr := `{"Name":"张三","Age":25,"Rmb":12345.67,"Hobby":["抽烟","喝酒","赌博"],"Sex":true}`
	dateMap := make(map[string]interface{})
	bytes := []byte(jsonStr)
	err := json.Unmarshal(bytes, &dateMap)
	if err != nil {
		fmt.Println("转化错误")
	} else {
		fmt.Println(dateMap)
	}
}

/**
json转结构体object 第一种类型
{"Name":"张三","Age":25,"Rmb":12345.67,"Hobby":["抽烟","喝酒","赌博"],"Sex":true}
*/
func main2() {
	var jsonStr = `{"Name":"张三","Age":25,"Rmb":12345.67,"Hobby":["抽烟","喝酒","赌博"],"Sex":true}`
	type Student struct {
		Name  string
		Age   int
		Rmb   float64
		Hobby []string
		Sex   bool
	}
	bytes := []byte(jsonStr)
	//student2 := &Student{}
	//等价
	student := new(Student)
	err := json.Unmarshal(bytes, student)
	if err != nil {
		fmt.Println("转化错误")
	} else {
		fmt.Println(*student)
	}
}

/**
json转切片，也就是没有标签名的数组 第二种类型
[{"age":25,"hobby":["抽烟","喝酒","赌博"],"name":"张三","rmb":123.45},
{"age":25,"hobby":["抽烟","喝酒","赌博"],"name":"张三","rmb":123.45},
{"age":25,"hobby":["抽烟","喝酒","赌博"],"name":"张三","rmb":123.45}]
*/
func main3() {
	jsonStr := `[{"age":25,"hobby":["抽烟","喝酒","赌博"],"name":"张三","rmb":123.45},
{"age":25,"hobby":["抽烟","喝酒","赌博"],"name":"张三","rmb":123.45},
{"age":25,"hobby":["抽烟","喝酒","赌博"],"name":"张三","rmb":123.45}]`
	dateSlice := make([]map[string]interface{}, 0)
	bytes := []byte(jsonStr)
	err := json.Unmarshal(bytes, &dateSlice)
	if err != nil {
		fmt.Println("转化错误")
	} else {
		fmt.Println(dateSlice)
	}
}

/**
切片转结构体
*/
func main() {
	jsonStr := `[{"age":25,"hobby":["抽烟","喝酒","赌博"],"name":"张三","rmb":123.45},
{"age":25,"hobby":["抽烟","喝酒","赌博"],"name":"张三","rmb":123.45},
{"age":25,"hobby":["抽烟","喝酒","赌博"],"name":"张三","rmb":123.45}]`
	type Student struct {
		Name  string
		Age   int
		Rmb   float64
		Hobby []string
		Sex   bool
	}
	students := make([]Student, 0)
	bytes := []byte(jsonStr)
	err := json.Unmarshal(bytes, &students)
	if err != nil {
		fmt.Println("转化错误")
	} else {
		fmt.Println(students)
	}
}
