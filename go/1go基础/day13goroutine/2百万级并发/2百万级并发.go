package main

import (
	"fmt"
	"time"
)

func DoTask(no int) {
	for i := 0; i < 10; i++ {
		fmt.Println(no)
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	//可以百万级并发
	for i := 0; i < 1000000; i++ {
		go DoTask(i)
	}
	// 主协程死，子协程跟着死
	fmt.Println("main over")
	//time.Sleep(6 * time.Second)
}
