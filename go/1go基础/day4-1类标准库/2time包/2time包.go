package main

import (
	"fmt"
	"time"
)

func main1() {
	nowTime := time.Now()
	fmt.Println(nowTime)

	fmt.Println(nowTime.Year())
	fmt.Println(nowTime.Month())
	fmt.Println(nowTime.Day())

	year, month, day := nowTime.Date()
	fmt.Println(year, month, day)

	nowTime.Day()
	nowTime.YearDay()
	nowTime.Weekday()

	nowTime.Hour()
	nowTime.Minute()
	nowTime.Second()
	nowTime.Nanosecond()

}

/**
创建时间
*/
func main2() {
	date := time.Date(-211, time.September, 8, 15, 15, 0, 0, time.Now().Location())
	fmt.Println(date)
}

/**
解析时间差
*/
func main() {
	now := time.Now()
	//一天之前
	duration, err := time.ParseDuration("-24h0m0s")
	if err != nil {
		fmt.Println(now.Add(duration))
	}

	//一周之前
	fmt.Println(now.Add(7 * duration))
	//一月之前
	fmt.Println(now.Add(30 * duration))

	//时间差，昨天和今天的时间的差
	sub := now.Sub(now.Add(duration))
	fmt.Println(sub)
}
