package main

import (
	"fmt"
	"sync"
	"time"
)

/*
//添加一个协程
wg.Add(1)
//减掉一个协程
wg.Done()
//等待组中协程数目为0时，不再等待
wg.Wait()
*/
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("A协程over")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		time.After(4 * time.Second)
		fmt.Println("B协程over")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for i := 0; i < 4; i++ {
			<-ticker.C
		}
		ticker.Stop()
		fmt.Println("C协程over")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("main over")
}
