package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
}

type Admin struct {
	User
	title string
}

/**
1、没有指针
2、如果是指针，应该先使用 Elem 方法获取目标类型，指针本身是没有字段成员的
 */
func main1() {
	var u Admin
	t := reflect.TypeOf(u)

	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}

/**
2、如果是指针，应该先使用 Elem 方法获取目标类型，指针本身是没有字段成员的
3、Type可以获取所有的属性字段
 */
func main() {
	u := new(Admin)

	t := reflect.TypeOf(u)
	fmt.Println(t.Kind())
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}