package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
	age int
}

type Admin struct {
	User
	title string
	Name string
}

/**
1、查找，基本类型
2、私有非导出类型
 */
func main1() {
	u := &Admin{User{"Jack", 23}, "NT","zhangsan"}
	v := reflect.ValueOf(u).Elem()

	fmt.Println(v.FieldByName("title").String()) // 用转换方法获取字段值
	fmt.Println(v.FieldByName("age").Int()) // 直接访问嵌入字段成员
	fmt.Println(v.FieldByIndex([]int{0, 1}).Int()) // 用多级序号访问嵌入字段成员

	//使用interface()只能获取Name值，不能获取name值,需要再强转为对应类型
	fmt.Println(v.FieldByName("Name").Interface().(string))
	name := v.FieldByName("Name")
	//函数Kind()是struct类型，属性值是interface（int byte等)
	fmt.Println(name.Type().String(),name.Kind())
}

/**
2除具体的 Int、String 等转换方法，还可返回 interface{}。只是非导出字段无法使用，需用 CanInterface 判断一下。
当然，转换成具体类型不会引发 panic。
 */
func main() {
	u := User{"Jack", 23}
	v := reflect.ValueOf(u)
	fmt.Println(v.FieldByName("Username").Interface())
	//age私有报错,使用之前判断一下
	if v.FieldByName("age").CanInterface() {
		fmt.Println(v.FieldByName("age").Interface())
	} else {
		fmt.Println(v.FieldByName("age").Int())
	}
}