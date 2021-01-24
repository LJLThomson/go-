package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
1、函数传递，将函数当做对象传递
*/
func main() {
	fmt.Println(eval(4, 4, "*"))
	q, r := div(4, 5)
	fmt.Println(q, r)
	//_表示跳过
	_, r = div(4, 6)
	fmt.Println(r)
	if resutl, err := eval2(3, 5, "x"); err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Print(resutl)
	}

	fmt.Println(apply(pow, 3, 5))
	//匿名函数编程
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(3, 4, 5, 6, 7))

}
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "%":
		q, _ := div(a, b)
		return q
	default:
		panic("unsinged opration" + op)
	}
}

func div(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return q, r
	//return a / b, a % b
}

/**
改造
*/
func eval2(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	case "%":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("%s\n", op)
	}
}

/**
函数编程
*/
func apply(op func(int, int) int, a,  c int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Print("Calling func %s with args"+"(%d %d)", opName, a, c)
	return op(a, c)
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

/**
可变参数列表
*/
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s = numbers[i]
	}
	return s
}
