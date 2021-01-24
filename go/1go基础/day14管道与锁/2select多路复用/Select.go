package main

import (
	"fmt"
	"time"
)

/**
select 多路复用，随机选择
*/
func main1() {
	//写5次
	channelA := make(chan int, 5)
	//读三次
	channelB := make(chan int, 3)
	channelB <- 1
	channelB <- 2
	channelB <- 3
	//读四次
	channelC := make(chan int, 4)
	channelC <- 4
	channelC <- 5
	channelC <- 6
	channelC <- 7
	go TaskA(channelA)
	go TaskB(channelB)
	go TaskC(channelC)
	time.Sleep(10 * time.Second)
	fmt.Println("main over")
}

func main() {
	//写5次
	channelA := make(chan int, 5)
	//读三次
	channelB := make(chan int, 3)
	channelB <- 1
	channelB <- 2
	channelB <- 3
	//读四次
	channelC := make(chan int, 4)
	channelC <- 4
	channelC <- 5
	channelC <- 6
	channelC <- 7
OUTER:
	for {
		//随机选择一路调用,直到所有路（管道）都阻塞时，执行default跳出
		select {
		case channelA <- 123:
			time.Sleep(time.Second)
			fmt.Println("执行A任务")
		case <-channelB:
			time.Sleep(time.Second)
			fmt.Println("执行B任务")
		case <-channelC:
			time.Sleep(time.Second)
			fmt.Println("执行C任务")
		default:
			//只能跳出select,如果需要跳出for,需要标签
			//break
			break OUTER
		}
	}
	fmt.Println("main over")
}

func TaskA(channel chan int) {
	for {
		channel <- 123
		time.Sleep(1 * time.Second)
	}
}

func TaskB(channel chan int) {
	for {
		<-channel
		time.Sleep(1 * time.Second)
	}
}

func TaskC(channel chan int) {
	for {
		<-channel
		time.Sleep(1 * time.Second)
	}
}
