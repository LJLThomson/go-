package main

import (
	"fmt"
	"os"
	"time"
)

/**
终止
*/
func main1() {
	timer := time.NewTimer(3 * time.Second)
	go func() {
		time.Sleep(2 * time.Second)
		ok := timer.Stop()
		//关闭不了输出型管道，只能是写入型
		//close(timer.C)
		//程序退出
		if ok {
			os.Exit(0)
		}
	}()
	x := <-timer.C
	fmt.Println("爆炸时间点", x)
}

/**
重置
*/
func main() {
	timer := time.NewTimer(5 * time.Second)
	//重置为2秒，从当前时间开始计算
	reset := timer.Reset(2 * time.Second)
	fmt.Println("重置成功", reset)
	x := <-timer.C
	fmt.Println("爆炸时间点", x)
}
