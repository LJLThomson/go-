package main

import (
	"fmt"
	"project2/day4interface"
)

/**
1、将Queue改造成任意类型
2、限定特定类型
*/
func main() {
	//将
	q := day4interface.Queue{1}
	q.Push(2)
	q.Push(3)
	q.Push("abc")
	fmt.Println(q.Pop())
	fmt.Print(q.Pop())
	fmt.Print(q.IsEmpty())

	myQueue := day4interface.MyQueue{1}
	myQueue.Push(2)
	myQueue.Push(3)
	//报错
	//myQueue.Push("abc")
	fmt.Println(myQueue.Pop())
	fmt.Print(myQueue.Pop())
	fmt.Print(myQueue.IsEmpty())
}
