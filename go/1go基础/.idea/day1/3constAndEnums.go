package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	consts()
	enums()
	increEnums()
	moveEnums()
}

/**
常量
*/
func consts() {
	const c, d = 5, 6
	const (
		a, b     = 3, 4
		filename = "abc"
	)
	fmt.Println(cmplx.Sqrt(a*a + b*b))
	var f int
	f = int(math.Sqrt(a*a + b*b))
	fmt.Println(f)
}

/**
枚举
*/
func enums() {
	const (
		cpp        = 0
		java       = 1
		python     = 2
		javascript = 3
	)
	fmt.Println(cpp, java, python, javascript)
}

/**
自增枚举
*/
func increEnums() {
	const (
		cpp = iota
		_
		java
		python
		javascript
	)
	fmt.Println(cpp, java, python, javascript)
}

/**
左右移自增
*/
func moveEnums() {
	const (
		b = 1 << (10 * iota)
		kb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, gb, tb, pb)
}
