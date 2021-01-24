package main

import (
	"fmt"
	"project2/day4interface"
)

/**
1、 duck typing 描述外部的行为而非内部结构，interface,注释来说明interface方法
2、 download ———— retrive
3、接口是隐式的,类似于java中工厂模式
*/
func main() {
	var r day4interface.IRetriver
	//1值类型
	r = day4interface.Retriver{}
	if retriver, ok := r.(day4interface.Retriver);ok {
		fmt.Println(retriver)
	} else {
		fmt.Println("转化失败")
	}

	//fmt.Println(download(r))
}

func download(r day4interface.IRetriver) string {
	return r.Get("http://wwww.baidu.com")
}
