package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	Username string
	age int
}
/**
将对象转换为接口，会发生复制行为。该复制品只读，无法被修改。所以要通过接口改变目标对象状态，必须是 pointer-interface。
就算是指针，我们依然没法将这个存储在 data 的指针指向其他对象，只能透过它修改目标对象。因为目标对象并没有被复制，被复制的只是指针。
总结，无法通过对象修改私有值，可以通过指针来操作
 */
func main1() {
	u := User{"Jack", 23}

	v := reflect.ValueOf(u)
	p := reflect.ValueOf(&u)

	fmt.Println(v.CanSet(), v.FieldByName("Username").CanSet())
	fmt.Println(p.CanSet(), p.Elem().FieldByName("Username").CanSet())
}

/**
1、私有字段修改
 */
func main2() {
	u := User{"Jack", 23}
	p := reflect.ValueOf(&u).Elem()
	//公共可以直接设置
	p.FieldByName("Username").SetString("Tom")

	//私有通过指针
	f := p.FieldByName("age")
	fmt.Println(f.CanSet())

	// 判断是否能获取地址。
	if f.CanAddr() {
		age := (*int)(unsafe.Pointer(f.UnsafeAddr()))
		// age := (*int)(unsafe.Pointer(f.Addr().Pointer())) // 等同
		*age = 88
	}

	// 注意 p 是 Value 类型，需要还原成接口才能转型。
	fmt.Println(u, p.Interface().(User))
}

/**
复合类型map slice等修改字段
 */
func main() {
	s := make([]int, 0, 10)
	v := reflect.ValueOf(&s).Elem()

	v.SetLen(2)
	//修改字段
	v.Index(0).SetInt(100)
	v.Index(1).SetInt(200)

	fmt.Println(v.Interface(), s)

	//添加字段
	v2 := reflect.Append(v, reflect.ValueOf(300))
	v2 = reflect.AppendSlice(v2, reflect.ValueOf([]int{400, 500}))

	fmt.Println(v2.Interface())

	fmt.Println("----------------------")

	m := map[string]int{"a": 1}
	v = reflect.ValueOf(&m).Elem()

	v.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(100)) // update
	v.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(200)) // add
	fmt.Println(v.Interface(), m)
}