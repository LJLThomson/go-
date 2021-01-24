package main

import (
	"fmt"
	"sync"
	"time"
)

/**
同步锁
mt sync.Mutex
mt.Lock()
mt.Unlock()
*/
func main() {
	var mt sync.Mutex
	var wg sync.WaitGroup
	var money = 2000
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			fmt.Println("协程%d准备抢锁", index)
			mt.Lock()
			fmt.Println("协程%d抢到锁", index)
			for j := 0; j < 10000; j++ {
				money += 1
			}
			<-time.After(3 * time.Second)
			mt.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("main over")
}
