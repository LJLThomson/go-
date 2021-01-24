package main

import "fmt"

/**
cooperation ---合作
routine ---事务
co-routine
goroutine go程
*/
func main1() {
	//go主协程
	CountNumber1()
	CountNumber2()
}

func main() {
	//go子协程
	go CountNumber1()
	//go主协程
	CountNumber2()
}

func CountNumber1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
func CountNumber2() {
	for i := 91; i < 100; i++ {
		fmt.Println(i)
	}
}
