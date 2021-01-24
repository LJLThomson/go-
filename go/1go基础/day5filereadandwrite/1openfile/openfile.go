package main

import (
	"fmt"
	"os"
	"time"
)

/**
1、文件打开
2、文件关闭
*/
func main() {
	file, err := os.Open("abc.txt")
	if err != nil {
		fmt.Println("文件出错，查找不到改文件")
		return
	}
	defer func() {
		file.Close()
		fmt.Println("文件已关闭")
	}()
	fmt.Println(file)
	time.Sleep(1 * time.Second)
}
