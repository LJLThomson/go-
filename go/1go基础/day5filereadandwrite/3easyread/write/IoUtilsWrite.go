package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dateStr := `想当年
    老夫拳打南山敬老院
	脚踢北海幼儿园
		我骄傲了吗`
	bytes := []byte(dateStr)
	err := ioutil.WriteFile("c.txt", bytes, 0666)
	if err != nil {
		fmt.Println("打开失败")
	} else {
		fmt.Println("打开成功")
	}
}
