package main

import (
	"fmt"
	"math"
)

type invlideRadiusError struct {
	Radius    float64
	MinRadius float64
	MaxRadius float64
}

func (i *invlideRadiusError) Error() string {
	info := fmt.Sprintf("%f当前半径，最大半径%f,最小半径%f", i.Radius, i.MaxRadius, i.MinRadius)
	return info
}

//2温和手段，error
func GetToyBollVolumn2(radius float64) (volumn float64, err error) {
	if radius < 0 {
		panic(&invlideRadiusError{radius, 5, 20})
	}
	if radius < 5 || radius > 20 {
		err = &invlideRadiusError{radius, 5, 20}
		//会返回默认值，float64默认值
		return
	}
	return (4 / 3.0) * math.Pi * math.Pow(radius, 3), nil
}

/**
自定义异常，实现error接口，这样本身就是一个error,类似于errors.go文件
*/
func main() {
	volumn2, err := GetToyBollVolumn2(3)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(volumn2)
	}
}
