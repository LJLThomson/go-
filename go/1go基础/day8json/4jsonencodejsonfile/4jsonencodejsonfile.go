package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/**
1、解码，将json文件数据往object对象或slice中
2、解码为map
3、解码时文件只可以为可读，不可以是os.O_RDONLY|os.O_WRONLY
*/
func main1() {
	srcFile, err := os.OpenFile("d.json", os.O_RDONLY, 0666)
	//srcFile, err := os.Open("d.json")
	defer func() {
		srcFile.Close()
	}()
	if err != nil {
		fmt.Println("打开文件错误")
		return
	}
	decoder := json.NewDecoder(srcFile)
	dateMap := make(map[string]interface{})
	err = decoder.Decode(&dateMap)
	if err != nil {
		fmt.Println("解码错误")
	} else {
		fmt.Println("解码成功")
		fmt.Println(dateMap)
	}
}

/**
解码为结构体
*/
func main() {
	srcFile, err := os.OpenFile("e.txt", os.O_RDONLY, 0666)
	defer func() {
		srcFile.Close()
	}()
	if err != nil {
		fmt.Println("打开文件错误")
		return
	}
	decoder := json.NewDecoder(srcFile)
	//其中Student结构体中Name|Age|Rmb注意大小写
	type Student struct {
		Name  string
		Age   int
		Rmb   float64
		Hobby []string
		Sex   bool
	}
	students := make([]Student, 0)
	err = decoder.Decode(&students)
	if err != nil {
		fmt.Println("解码错误")
	} else {
		fmt.Println("解码成功")
		fmt.Println(students)
	}
}
