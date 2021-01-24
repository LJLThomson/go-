package main

import "fmt"

/**
1、go指针不能加1，不同于C/C++中的指针
2、go是采用值传递（拷贝），java/python是引用传递（不拷贝）(基本语言int\bool是值传递）,c/c++有值传递也有引用传递
*/
func main() {
	//1值传递
	a, b := 3, 4
	swap(a, b)
	//2指针传递
	fmt.Println(a, b)
	funcSwap(&a, &b)
	fmt.Println(a, b)
	c, d := valueSwap(3, 4)
	fmt.Println(c, d)
}

func swap(a, b int) {
	a, b = b, a
}

func funcSwap(a, b *int) {
	*a, *b = *b, *a
}

func valueSwap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}
