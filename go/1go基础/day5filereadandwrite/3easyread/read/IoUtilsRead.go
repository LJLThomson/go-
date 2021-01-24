package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//文件自动关闭
	bytes, err := ioutil.ReadFile("abc.txt")
	if err != nil {
		fmt.Println("读取失败")
	} else {
		contentStr := string(bytes)
		fmt.Println(contentStr)
	}
}
