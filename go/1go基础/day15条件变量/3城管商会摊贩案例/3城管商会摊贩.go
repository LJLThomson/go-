package main

import (
	"fmt"
	"sync"
	"time"
)

/**
监听城管大队
烧烤摊集群：监听城管大队，只要出动就进行阻塞等待至唤醒，否则就提供露天烧烤
工会：公关城管，并广播通知烧烤摊

1、城管来了，摊贩进行收摊，发出城管预警
2、城管来了，工会进行打通关系
3、工会打通关系，摊贩继续贩卖
*/
func main() {
	cond := sync.NewCond(&sync.Mutex{})
	var dangerous = false
	channel := make(chan string,1)
	for i := 0; i < 3; i++ {
		go func(index int) {
			for {
				cond.L.Lock()
				for dangerous == true {
					select {
					case channel <- "危险，城管来了":
						fmt.Println(index, "城管来了")
					default: //有人已经发送了
					}
					//进行等待,进行阻塞，直到收到广播通知
					cond.Wait()
				}
				cond.L.Unlock()
				//没有城管，可以提供露天烧烤了
				fmt.Println(index, "没有城管，可以提供烧烤")
				<-time.After(3 * time.Second)
			}
		}(i)
	}

	go func() {
		for {
			select {
			case <-channel:
				cond.L.Lock()
				//工会花了五天时间处理城管
				<-time.After(5 * time.Second)
				dangerous = false
				fmt.Println("危险已经摆平")
				cond.Broadcast()
				cond.L.Unlock()
			default:
				fmt.Println("工会过自己幸福休闲的日子直到城管出动了")
				time.Sleep(6 * time.Second)
				dangerous = true
			}
		}
	}()

	time.Sleep(365 * time.Second)
}
