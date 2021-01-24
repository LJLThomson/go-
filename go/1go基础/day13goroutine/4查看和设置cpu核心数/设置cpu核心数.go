package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("当前协程的可用cpu逻辑核心数为：", runtime.NumCPU()) //本电脑为4，不会变
	//max processors
	//将可用cpu核心数设置为4，返回原来值核心数（本电脑4）
	fmt.Println("当前协程可用的核心数", runtime.GOMAXPROCS(2)) // 4
	fmt.Println("当前协程可用的核心数", runtime.GOMAXPROCS(1)) //2
}
