package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	f2 := fibonacci2()
	printFileContents(f2)
}

/**
斐波拉契数列
*/
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

/**
数列读取存储
*/
type intGen func() int

func fibonacci2() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

/**
实现读取的接口
p byte
*/
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	// TODO:incorrect if p is too small
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	// scanner.Scan进行读取
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
