package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
	age int
}

func main1() {
	u := User{}
	v := reflect.ValueOf(u)

	f := v.FieldByName("a")
	//查找不到则f.kind()为invalid无效
	fmt.Println(f.Kind(), f.IsValid())
}

func main() {
	var p *int

	var x = p
	fmt.Println(x == nil)

	v := reflect.ValueOf(p)
	fmt.Println(v.Kind(), v.IsNil())


}