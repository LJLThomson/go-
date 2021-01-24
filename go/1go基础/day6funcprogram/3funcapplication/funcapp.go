package main

import (
	"fmt"
	"project2/day6funcprogram"
)

/**
1、遍历二叉树，先左边数，后右边数，最后中间数，在java中实现起来很复杂， 但在go\c++中很简单
2、为函数实现接口，也就是传函数到参数中，函数作参数，函数中自己有方法
		0
	2 		 3
4				5
*/
func main() {
	mynode := day6funcprogram.MyNode{Value: 0}
	mynode.Left = &day6funcprogram.MyNode{Value: 2}
	mynode.Left.Left = &day6funcprogram.MyNode{Value: 4}
	mynode.Right = &day6funcprogram.MyNode{Value: 3}
	mynode.Right.Right = &day6funcprogram.MyNode{Value: 5}
	mynode.Travese()

	numCount := 0
	//传不同函数进去，就可以调用不同方法，类似java中匿名内部类，或者说接口，比如监听方法.setOnClicker(new ...)
	mynode.TraverseFunc(func(node *day6funcprogram.MyNode) {
		numCount++
	})
	fmt.Println("Node count:", numCount)
}
