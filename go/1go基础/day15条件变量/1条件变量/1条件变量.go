package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	bitcoinRasing := false
	go func() {
		for {
			time.Sleep(3 * time.Second)
			cond.L.Lock()
			bitcoinRasing = true
			//cond.Signal()  只通知一个正在wait的条件
			cond.Broadcast() // 通知所有正在wait的条件
			cond.L.Unlock()

			//3秒之后开始跌
			<-time.After(3 * time.Second)
			cond.L.Lock()
			bitcoinRasing = false
			cond.Signal()
			cond.L.Unlock()
			<-time.After(1 * time.Second)
		}
	}()
	for {
		cond.L.Lock()
		for !bitcoinRasing {
			fmt.Println("停止投资")
			cond.Wait() //阻塞，并释放锁，监听到之后，再次加锁
		}
		fmt.Println("开始投资")
		cond.L.Unlock()
		<-time.After(1 * time.Second)
	}

}
