package main

import (
	"fmt"
	"sync"
	"time"
)

/**
sync.RWMutex
	//读写锁
	rwm.RLock()
	rwm.Unlock()
	rwm.Lock()
	rwm.Unlock()
读时，多路可读
写时，一路可写
读写互斥
*/
func main() {
	var rwm sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		rwm.RLock()
		fmt.Println("读数据库")
		<-time.After(3 * time.Second)
		rwm.RUnlock()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		rwm.RLock()
		fmt.Println("读数据库")
		<-time.After(3 * time.Second)
		rwm.RUnlock()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		rwm.RLock()
		fmt.Println("读数据库")
		<-time.After(3 * time.Second)
		rwm.RUnlock()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		rwm.RLock()
		fmt.Println("读数据库")
		<-time.After(3 * time.Second)
		rwm.RUnlock()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		rwm.Lock()
		fmt.Println("写入数据库")
		<-time.After(3 * time.Second)
		rwm.Unlock()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		rwm.Lock()
		fmt.Println("写入数据库")
		<-time.After(3 * time.Second)
		rwm.Unlock()
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		rwm.Lock()
		fmt.Println("写入数据库")
		<-time.After(3 * time.Second)
		rwm.Unlock()
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("main over")
}
