package main

import (
	"fmt"
	"math"
)

/**
创建工厂模式
*/
type invlideRadiusError struct {
	Radius    float64
	MinRadius float64
	MaxRadius float64
}

func (i *invlideRadiusError) Error() string {
	info := fmt.Sprintf("%f当前半径，最大半径%f,最小半径%f", i.Radius, i.MaxRadius, i.MinRadius)
	return info
}

/**
工厂模式返回一个对象指针
*/
func NewInvolidRadiusError(radius float64) *invlideRadiusError {
	ire := new(invlideRadiusError)
	ire.MinRadius = 5
	ire.MaxRadius = 20
	ire.Radius = radius
	return ire
}

//2温和手段，error
func GetToyBollVolumn2(radius float64) (volumn float64, err error) {
	if radius < 0 {
		panic(NewInvolidRadiusError(radius))
	}
	if radius < 5 || radius > 20 {
		err = NewInvolidRadiusError(radius)
		//会返回默认值，float64默认值
		return
	}
	return (4 / 3.0) * math.Pi * math.Pow(radius, 3), nil
}

func main() {
	defer func() {
		// 异常处理机制，不让程序崩溃
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	volumn2, err := GetToyBollVolumn2(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(volumn2)
	}
}
