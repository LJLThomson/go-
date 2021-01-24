package main

import "fmt"

/**
1、基本数组定义
2、遍历数组
3、数组传递，go可以传， c++中不能值传递，可以引用或指针传递
4、go语言中一般不适用数组，就像java中使用list代替数组，C++中使用vector
*/
func main() {
	//1赋初值
	var array1 [5]int
	//:需要赋初值
	array2 := [3]int{1, 2, 3}
	//定义多个初始值
	array3 := [...]int{2, 4, 6, 8, 10}
	var grad [4][5]int
	fmt.Println(array1, array2, array3)
	fmt.Println(grad)
	//2遍历数组
	for i := 0; i < len(array3); i++ {
		fmt.Println(array3[i])
	}
	//同时获取下标和值
	//java和python只有foreach
	for i, v := range array3 {
		fmt.Print(i, v)
	}
	for i, _ := range array3 {
		fmt.Println(i)
	}
	sum := 0
	for _, v := range array3 {
		sum += v
	}
	fmt.Println(sum)
	array4 := [5]int{2, 3, 4, 5, 6}
	//这里C++代表首地址指针，go中代表数组
	printArray(array4)
	//array4值不会变，值传递
	fmt.Println(array4)
	printArray2(&array4)
}

/**
3拷贝方式
*/
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

/**
3指针方式
go中参数指针,形参写法只有类似*[5]int,没有*arr [5]int
*/
func printArray2(arr *[5]int) {
	//C++中写法，go中可以直接arr[0]也能找到
	//(*arr)[0] = 100
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
