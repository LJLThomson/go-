package _定时器

import (
	"fmt"
	"time"
)

func main1() {
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("计时开始", time.Now())
	//采用阻塞的方式停止等待
	x := <-timer.C
	fmt.Println("引爆时间到", x)
}

/**
简易方式
 */
func main2() {
	//封装
	x := <-time.After(3 * time.Second)
	fmt.Println(x)
}

