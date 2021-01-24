package day3Object

import "fmt"

type Student struct {
	Age    int
	Name   string
	Height int
}

func (stu *Student) Print() {
	fmt.Println(stu.Age)
}

func (stu *Student) SetAge(age int) {
	stu.Age = age
}
