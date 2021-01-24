package main

import (
	"fmt"
	"time"
)

/**
1、和map类似，一片内存空间，传参时都是引用，指向同一个对象
2、make方式指向一片内存空间
3、size缓存区最大长度
*/
func main1() {
	//1、声明
	//var channel chan int
	//fmt.Println(channel)

	//2
	//channel := make(chan int)
	//fmt.Println(channel)

	//3size最大容量
	channel := make(chan int, 3)
	fmt.Println(channel)
}

/**
1、chan管道中缓存长度不够时，写不进去，主协程永远阻塞，类似于IO流缓冲区
*/
func main2() {
	//1
	//channel := make(chan int)
	//1 修正，可以写三个值
	channel := make(chan int, 3)
	//写入
	channel <- 234
	channel <- 235
	channel <- 236
	//读取
	x := <-channel
	fmt.Println(x)
	fmt.Println("main gameover")
}

/**
1、读写细节，写入，管道满了无法写入，主协程等待
2、读取，管道空了之后，子线程等待
*/
func main3() {
	channel := make(chan int, 3)

	go func() {
		for i := 0; i < 5; i++ {
			// 已经读空，没有可读，就阻塞子线程读取位置
			time.Sleep(3 * time.Second)
			x := <-channel
			fmt.Println(x)
		}
	}()

	for i := 0; i < 4; i++ {
		//写入管道满了之后，没有人读，就阻塞在写入状态
		//3时满了，写入等待
		channel <- i
		fmt.Println("channel game over")
	}
	for {
		time.Sleep(3 * time.Second)
	}
}

/**
close
1、无法向一个已经关闭的管道中写入数据，但可以读取
2、判断当前管道情况
3、不能关闭未初始化的管道，var channel chan int
4、不能重复关闭管道
*/
func main4() {
	channel := make(chan int, 3)
	go func() {
		time.Sleep(4 * time.Second)
		x := <-channel
		fmt.Println(x)
		x = <-channel
		fmt.Println(x)
		//读取的x为0，默认值，不是有效值
		x = <-channel
		fmt.Println(x)

		//2、当前管道情况，false管道关闭，true,关闭打开状态
		x, ok := <-channel
		if ok {
			fmt.Println("能读到有效值", x)
		} else {
			fmt.Println("不能读到有效值")
		}
	}()

	channel <- 1
	channel <- 2
	//1、close，可以读取，不能再次写入
	close(channel)
	for {
		time.Sleep(3 * time.Second)
	}
}

/**
range遍历读取
1、管道未关闭时，会一直读取，会阻塞
2、close管道，会广播形式通知所有读取的子协程不再阻塞（不会再写入了），子协程走下一步
*/
func main() {
	channel := make(chan int, 3)
	go func() {
		time.Sleep(4 * time.Second)
		//1管道未关闭时，会一直读取，会阻塞
		for x := range channel {
			fmt.Println(x)
		}
		fmt.Println("game over")
	}()

	channel <- 1
	channel <- 2
	//2、close管道，会广播形式通知所有读取管道数据的协程读取不再阻塞
	close(channel)
	for {
		time.Sleep(3 * time.Second)
	}
}
