package main

import (
	"fmt"
	"sync"
)

/**
发生在主协程，不会报错，但在子协程依旧属于死锁
AB相互都要求对方先发红包，自己再发
*/
func main1() {
	var wg sync.WaitGroup
	chanA := make(chan string, 3)
	chanB := make(chan string, 4)
	wg.Add(1)
	go func() {
		<-chanA
		chanB <- "da"
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		<-chanB
		chanA <- "xx"
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("main over")
}

/**
读写相互锁定对方需要的资源（当管道中缓存区不够时）
*/
func main() {
	var rwm sync.RWMutex
	var wg sync.WaitGroup
	chanA := make(chan int)
	wg.Add(1)
	go func() {
		rwm.RLock()
		x := <-chanA
		fmt.Println("读到", x)
		rwm.RUnlock()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		rwm.Lock()
		chanA <- 123
		rwm.Unlock()
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("main over")
}
