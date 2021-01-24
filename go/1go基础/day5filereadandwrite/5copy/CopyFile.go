package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, _ := ioutil.ReadFile("abc.txt")
	err := ioutil.WriteFile("c.txt", bytes, 0666)
	if err != nil {
		fmt.Println("读取失败")
	} else {
		contentStr := string(bytes)
		fmt.Println(contentStr)
	}
}
