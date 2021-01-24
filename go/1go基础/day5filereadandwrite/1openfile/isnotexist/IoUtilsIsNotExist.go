package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("d.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
		}
	} else {
		fmt.Println("文件存在", fileInfo)
	}
}
