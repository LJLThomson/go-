package main

import (
	"fmt"
	"os"
)

// 获取命令行参数，用的不多
func main() {
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}
