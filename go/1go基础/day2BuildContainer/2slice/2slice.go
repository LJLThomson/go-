package main

import "fmt"

/**
1、切片slice概念，半开半关闭区间
2、slice是对原数据的一个view,相当于*arr指针
3、reslice,再次切片
4、slice扩展（取不到），三个值，ptr指针，len,cap,可以向后扩展，但不能超越底层数组，不能向前扩展
*/
func main() {
	//1、slic概念
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//fmt.Println("arr[2:6] = ", arr[2:6])
	//fmt.Println("arr[:6] = ", arr[:6])
	//fmt.Println("arr[2:] = ", arr[2:])
	//fmt.Println("arr[:] = ", arr[:])
	//2、对原数据一个view
	s1 := arr[1:5]
	updateArrayValue(s1)
	fmt.Println(s1)
	fmt.Println(arr)
	//3、
	s2 := s1[2:4]
	fmt.Println(s2)
	//4slice扩展,超越s2范围,但不超过底层数组arr,s2[5]这种写法时错误的
	s3 := s2[:5]
	fmt.Println(s3)
}

/**
[]int代表slice
*/
func updateArrayValue(arr []int) {
	arr[0] = 100
}
