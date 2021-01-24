package main

import (
	"fmt"
	"project2/day4interface"
)

func main() {
	var iretrive day4interface.IRetriver
	iretrive = myInterface{"lisi", 20}
	iretrive.Get("www")

	var combineInter day4interface.InterfaceC
	combineInter = myInterface{"zhangsan", 25}
	combineInter.Get("dafsdfs")
	combineInter.Post()
}

type myInterface struct {
	Name string
	Age  int
}

func (m myInterface) Get(url string) string {
	fmt.Println(url)
	return url
}

func (m myInterface) Post() {
	fmt.Println("post")
}

func (m myInterface) Read() {
	fmt.Println("read")
}

func (m myInterface) Write() {
	fmt.Println("write")
}
