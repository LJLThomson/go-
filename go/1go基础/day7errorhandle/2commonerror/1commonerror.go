package main

import "fmt"

/**
1、空指针异常
2、map下标越界
 */
func main() {
	//1空指针异常
	//var intptr *int
	//fmt.Println(*intptr)

	//2、slice下标越界
	//slice2 := []int{1,2,3}等价于下面
	//slice1 := make([]int, 0)
	//slice1 = append(slice1, 1, 2, 3, 4)
	//fmt.Println(slice1[5])
	//var varmap map[string]interface{}

	//map make就已经付初值了
	//m := make(map[string]interface{})
	//m["我的"] = 2
	//m["学生"] = 3
	//fmt.Println(m)

	//map 空值
	var varmap map[string]interface{}
	varmap["xx"] = 32
	fmt.Println(varmap)
}
