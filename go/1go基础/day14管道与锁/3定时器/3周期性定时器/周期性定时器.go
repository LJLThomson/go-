package main

import (
	"fmt"
	"time"
)

/**
每个一段时间触发一次
*/
func main() {
	//每隔一秒触发一次
	ticker := time.NewTicker(1 * time.Second)
	flag := false
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		flag = true
	}()
	for {
		if false {
			//当ticker.stop时，这里读取会发生阻塞
			fmt.Println(<-ticker.C)
		}
	}
}
