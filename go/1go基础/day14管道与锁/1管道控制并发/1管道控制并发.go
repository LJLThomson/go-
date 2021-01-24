package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

/**
实现类似java中线程池，一次最多只能执行n个协程
管道满了之后，写不进去就会进行阻塞，直到能读取为止
1、先需要注册，注册一个名字，
2、任务完毕之后从管道中读出一个，为后来的
*/
func main() {
	//信号量，控制管道数目最大为5，并发数最大为5
	semaphore := make(chan string, 5)
	for i := 0; i < 100; i++ {
		go GetSqrt("管道数："+strconv.Itoa(i), i, semaphore)
	}
	for {
		time.Sleep(2 * time.Second)
	}
}

func GetSqrt(name string, n int, channel chan string) {
	//1、先需要注册，注册一个名字，写不进去就会进行阻塞，直到能读取为止
	channel <- name
	sqrt := math.Sqrt(float64(n))
	fmt.Println(sqrt)
	time.Sleep(3 * time.Second)
	//做完之后，从管道中读出一个
	<-channel
}
