package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("我的")
	}()
	//defer func(int, string) (int,string){
	//
	//}()
	//匿名函数去掉名字，作为方法时需要知道参数名字，作为参数时不需要知道参数名字
	defer func(a int, name string) (int, string) {
		fmt.Println(a)
		fmt.Println(name)
		return a, name
	}(20, "zhangsan") //(),表示传参
}

func theme(a int, name string) (int, string) {
	fmt.Println(a)
	fmt.Println(name)
	return a, name
}
