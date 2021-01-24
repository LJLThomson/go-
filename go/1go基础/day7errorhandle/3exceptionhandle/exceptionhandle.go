package main

import (
	"errors"
	"fmt"
	"math"
)

func GetToyBollVolumn(radius float64) float64 {
	if radius < 0 {
		panic("半径不能为复数")
	}
	return (4 / 3.0) * math.Pi * math.Pow(radius, 3)
}

func main1() {
	defer func() {
		// 异常处理机制，不让程序崩溃
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	// 程序死在这里
	volumn := GetToyBollVolumn(-1.0)
	fmt.Println(volumn)
}

//2温和手段，error
func GetToyBollVolumn2(radius float64) (volumn float64, err error) {
	if radius < 0 {
		panic("半径不能为复数")
	}
	if radius > 5 || radius < 20 {
		err = errors.New("合法半径范围不是5到20")
		//会返回默认值，float64默认值
		return
	}
	return (4 / 3.0) * math.Pi * math.Pow(radius, 3), nil
}

func main() {
	volumn2, err := GetToyBollVolumn2(10)
	if err != nil {
		fmt.Println("小球体积错误 err:", err)
	} else {
		fmt.Println(volumn2)
	}
}
