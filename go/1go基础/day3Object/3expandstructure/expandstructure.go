package main

import (
	"fmt"
	"project2/day3Object"
	"reflect"
)

/**
1、组合
2、别名s
 */
func main() {
	myTreeNode := day3Object.MyTreeNode{}
	myTreeNode.PrintValue()

	q := day3Object.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Print(q.Pop())
	fmt.Print(q.IsEmpty())

	//定义一个方法
	//methods := func(t reflect.Type) {
	//	for i, n := 0, t.NumMethod(); i < n; i++ {
	//		m := t.Method(i)
	//		fmt.Println(m.Name)
	//	}
	//}

}

func reflectvalue(t reflect.Type)  {

}
/**
`field:"username" type:"nvarchar(20)"` 标签标记 ORM Model 属性。
 */
type User struct {
	Name string `field:"username" type:"nvarchar(20)"`
	Age int `field:"age" type:"tinyint"`
}