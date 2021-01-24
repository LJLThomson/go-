package main

import "fmt"

/**
写入缓存
*/
func main1() {
	channel := make(chan int, 3)
	channel <- 1
	fmt.Println("管道长度", len(channel))
	fmt.Println("管道容量", cap(channel))
	channel <- 1
	channel <- 2
	//已满的管道会阻塞写入
	channel <- 3
}

func main() {
	channel := make(chan int, 3)
	channel <- 1
	channel <- 2
	channel <- 3
	//读取
	x := <-channel
	<-channel
	<-channel
	fmt.Println(x)

	// 已空的管道，读取会阻塞，主协程，发生死锁错误
	<-channel
}
