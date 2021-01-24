package main

import "fmt"

/**
1、什么是函数式编程，参数、变量，返回值都可以是函数(类似于面向对象，使用方法）并且类似对象指针一样的链式编程（stringbuilder)，
	C++中是函数指针，方法不让返回，但go中其实也是当做类似指针来使用，
2、闭包概念 ,闭包，是将返回的函数全部概括，包括其中自由变量，自由变量类似于GC根节点方式链接起来
		闭包
	函数
局部变量  自由变量
			 ....
3、正统函数：不可变性，不能有状态，只有常量和函数，只有一个参数，但在go中不适用
*/
func main() {
	a := adder()
	//a1与a是隔离的，不相干，其中自由变量也不相干，可以独立保存
	//a1 := adder()
	//类似于面向对象，但调用时用方法，比如a(i)，a函数类似对象指针，该对象改变了方法中值sum就会保存起来（自由变量）
	for i := 0; i < 10; i++ {
		fmt.Printf("1+2...+i = %d\n", a(i))
	}
	b := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, b = b(i)  //返回b,方法类似于C++函数指针返回
		fmt.Printf("1+2...+i = %d\n", i, s)
	}
}

/**
函数编程，类似于一种对象指针，C++中函数指针很相似，闭包
*/
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

/**
正统函数编程,一种递归, 递归，重新adder2(base + v)获取b,方法相当于指针来使用了
*/

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (i int, i2 iAdder) {
		return base + v, adder2(base + v)
	}
}
