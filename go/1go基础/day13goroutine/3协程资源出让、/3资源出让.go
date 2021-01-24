package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
GoSche
*/
func main1() {
	for i := 0; i < 2; i++ {
		/**
		协程内的代码执行在子协程
		协程有一定初始化时间,3
		*/
		go func() {
			for j := 1; j < 10; j++ {
				fmt.Println("协程%d:%d\n", i, j)
			}
		}()
	}
	time.Sleep(2 * time.Second)
}

func main() {
	for i := 0; i < 10; i++ {
		/**
		协程内的代码执行在子协程
		协程有一定初始化时间
		*/
		go func(no int) {
			for j := 1; j < 10; j++ {
				if no == 5 {
					runtime.Gosched()
				}
				fmt.Println("协程%d:%d\n", no, j)
				time.Sleep(500 * time.Millisecond)
			}
		}(i)
	}
	time.Sleep(10 * time.Second)
}
