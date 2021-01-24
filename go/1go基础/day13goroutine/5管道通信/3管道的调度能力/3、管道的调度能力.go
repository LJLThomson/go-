package main

import (
	"fmt"
	"time"
)

func Count(name string, num int, chanQuite chan string) {
	for i := 0; i < num; i++ {
		fmt.Println(name, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name + "完毕")
	//完毕之后，通知主协程我的工作完成了
	chanQuite <- name + "已完毕"
}

/**
所有子协程完毕之后，主协程才关闭
1、通过管道读取阻塞的方式实现
子协程完毕之后，写入一个数据
主协程读取数据，直到数据读取完为止
*/
func main() {
	chanQuite := make(chan string, 3)
	go Count("daughter", 4, chanQuite)
	go Count("son", 5, chanQuite)
	Count("main", 2, chanQuite)

	for i := 0; i < 3; i++ {
		<-chanQuite
	}
	fmt.Println("gameover")
}
