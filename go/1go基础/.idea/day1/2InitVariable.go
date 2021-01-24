package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
内建变量类型.
*/
func main() {
	fmt.Print("内建变量类型")
	euler()
	triangle()
}

/**
欧拉复数公式 e^iPi + 1 = 0, 3+4i
*/
func euler() {
	a := 3 + 4i //不能写成3 + 4*i, i是复数，4*i，i是变量
	fmt.Println(a)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

/**
类型只有强制转换
*/
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
