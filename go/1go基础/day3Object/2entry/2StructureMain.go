package main

import (
	"fmt"
	"project2/day3Object"
)

/**
1、同一目录下，只有一个类
2、一个目录下，只有一个main
3、为结构定义的方法必须在同一个目录下，可以是不同文件
*/
func main() {
	stu  := day3Object.Student{}
	stu.SetAge(20)
	stu.Print()

	node := day3Object.Node{}
	fmt.Println(node.Value)
}
