package main

import "fmt"

/**
1、defer，保证函数在结束时发生
2、匿名函数
*/
func main() {
	deferMethod()

	defer func() {
		fmt.Println("匿名函数")
	}()
}

func deferMethod() {
	defer fmt.Println("我的")
	fmt.Println("不是")
	fmt.Println("没有")
}
