package main

import (
	"fmt"
	"sync/atomic"
)

/**
原子操作
多协程中，更改数据时，数据可能是不安全的，
解决方案：
1、从api上进行加锁
2、原子操作，仅仅使用了基本类型，底层硬件支持，保证数据一定会成功，并且不会被其他协程访问
原子操作主要是用于计算，其他情况用加锁方式
*/
func main01() {
	var a int64 = 12
	val := atomic.LoadInt64(&a)
	fmt.Println(val)

	atomic.StoreInt64(&a, 35)
	fmt.Println(a)

	//返回旧的值
	old := atomic.SwapInt64(&a, 64)
	fmt.Println(old, a)

	swapSuccess := atomic.CompareAndSwapInt64(&a, 64, 128)
	fmt.Println(swapSuccess, a)
}

func main2() {
	var a int64 = 12
	a = 32
	//比较失败，模拟了并发修改a数据，a不安全
	swapSuccess := atomic.CompareAndSwapInt64(&a, 12, 65)
	fmt.Println(swapSuccess, a)
}
