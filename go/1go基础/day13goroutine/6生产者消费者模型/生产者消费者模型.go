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

计数，数据向主协程告知，10个数之后退出
二、整点时生产者罢工，主协程退出
*/
func main() {
	chanShop := make(chan Product, 1)
	chanCount := make(chan int, 5)
	go ProduceShop(chanShop)
	go ConsumeShop(chanShop, chanCount)
	for i := 0; i < 10; i++ {
		<-chanCount
	}
	//主协程被kill掉了，子协程都陪着kill
	fmt.Println("game over")
}

/**
消费者
消费一个产品
并告知主协程，主协程计数
*/
func ConsumeShop(chanRead <-chan Product, chanCount chan<- int) {
	for {
		product := <-chanRead
		fmt.Println("商品被吃了" + product.Name)
		chanCount <- 1
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
		//每隔一段时间生产一个
		time.Sleep(1 * time.Second)
	}
}
