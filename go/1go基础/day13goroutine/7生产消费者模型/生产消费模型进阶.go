package main

import (
	"fmt"
	"strconv"
	"time"
)

type Product struct {
	Name string
}

/**
1、生产者每天生产一件产品
2、生产者每生产一件产品，消费者消费一件产品
3、消费10轮之后，主协程退出,计数

4、整点时生产者罢工，主协程退出
*/
var Godie = make(chan string, 3)
var isStrike = false

func main() {
	chanShop := make(chan Product, 1)
	chanCount := make(chan int, 5)
	go ProduceShop(chanShop)
	go ConsumeShop(chanShop, chanCount)
	go watcher()
	//阻塞命令，读到之后，主协程结束
	<-Godie
	//主协程被kill掉了，子协程都陪着kill
	fmt.Println("game over")
}

/**
消费者
消费一个产品
*/
func ConsumeShop(chanRead <-chan Product, chanCount chan<- int) {
	for {
		product := <-chanRead
		fmt.Println("商品被吃了" + product.Name)
	}
}
func watcher() {
	for {
		second := time.Now().Second()
		if second == 0 {
			isStrike = true
		}
		time.Sleep(2 * time.Second)
	}
}

/**
生产者
1、生产一件产品
2、产品写入管道，让消费者读取
*/
func ProduceShop(chanShopWrite chan<- Product) {
	for {
		// 生产一件产品
		product := Product{Name: "产品" + strconv.Itoa(time.Now().Second())}
		chanShopWrite <- product
		fmt.Println("生产者向管道输送了产品:" + product.Name)
		if isStrike {
			Godie <- "罢工了"
		}
		//每隔一段时间生产一个
		time.Sleep(1 * time.Second)
	}
}
