package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/**
1、编码，将object对象转化为json文件
2、解码，将json文件转化为结构体object
3、普通文件转化为json文件
*/
func main() {
	dateMap := make(map[string]interface{})
	dateMap["name"] = "张三"
	dateMap["age"] = 25
	dateMap["rmb"] = 123.45
	dateMap["hobby"] = []string{"抽烟", "喝酒", "赌博"}
	dateMap["sex"] = true
	dstFile, err := os.OpenFile("d.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer func() {
		dstFile.Close()
	}()
	if err != nil {
		fmt.Println("打开文件错误")
		return
	}
	encoder := json.NewEncoder(dstFile)
	err = encoder.Encode(dateMap)
	if err != nil {
		fmt.Println("编码错误")
	} else {
		fmt.Println("编码成功")
	}
}

/**
切片编码为json文件
*/
func main2() {
	type Student struct {
		Name  string
		Age   int
		Rmb   float64
		Hobby []string
		Sex   bool
	}
	student1 := Student{"张帧1", 25, 12367.89, []string{"抽小熊猫", "喝花酒", "烫花卷发"}, true}
	student2 := Student{"张帧2", 26, 12367.89, []string{"抽小熊猫", "喝花酒", "烫花卷发"}, true}
	student3 := Student{"张帧3", 27, 12367.89, []string{"抽小熊猫", "喝花酒", "烫花卷发"}, true}
	students := make([]Student, 0)
	students = append(students, student1, student2, student3)
	dstFile, err := os.OpenFile("e.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("打开文件错误")
		return
	}
	encoder := json.NewEncoder(dstFile)
	err = encoder.Encode(students)
	if err != nil {
		fmt.Println("编码错误")
	} else {
		fmt.Println("编码成功")
	}
}
