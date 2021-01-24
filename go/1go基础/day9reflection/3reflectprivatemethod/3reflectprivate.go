package main

import (
	"fmt"
	"project2/day9reflection"
	"reflect"
)

/**
1、在go中,私有类型，可以获取到属性值和方法
2、除具体的 Int、String 等转换方法，还可返回 interface{}。只是非导出字段无法使用，需用 CanInterface 判断一下。
*/

func main1() {
	person := day9reflection.Person{"张三", 25}
	famliy := day9reflection.NewFamily2("历史", 26, &person)
	pType := reflect.TypeOf(famliy)
	oType := pType.Elem()
	fmt.Println("打印对象类型名称")
	fmt.Println(pType.Name())
	fmt.Println(pType.Kind())
	fmt.Println("属性值，注意反射中格外注意，属性只能是对象.属性，不能是指针.属性")
	for i := 0; i < oType.NumField(); i++ {
		field := oType.Field(i)
		//对象名称和对象类型
		fmt.Println(field.Name)
		fmt.Println(field.Type)
	}
	fmt.Println("方法,指针与普通方法差异")
	//反射中，需要注意指针与普通类型，这里反射不能自行区分
	fmt.Println(oType.NumMethod()) // 0 找不到Family对象为
	fmt.Println(pType.NumMethod()) // 3 *Family
}

func main2() {
	person := day9reflection.Person{"张三", 25}
	famliy := day9reflection.NewFamily2("历史", 26, &person)
	pValue := reflect.ValueOf(famliy)
	oValue := pValue.Elem()
	fmt.Println("打印对象类型名称")
	fmt.Println(oValue.Kind())
	fmt.Println("打印属性值")
	oValue.NumField()
	for i := 0; i < oValue.NumField(); i++ {
		field := oValue.Field(i)
		////正射，field是反射值，不能使用
		//fmt.Println(field.Int())
		////获取type属性,类似于type
		fmt.Println(field.Type())
	}
}

type User struct {
	Username string
	age int
}

type Admin struct {
	User
	title string
}

func (*User) ToString() {}
func (Admin) test()     {}

func main3() {

	u := Admin{}

	methods := func(t reflect.Type) {
		for i, n := 0, t.NumMethod(); i < n; i++ {
			m := t.Method(i)
			fmt.Println(m.Name)
		}
	}

	fmt.Println("--- value interface ---")
	methods(reflect.TypeOf(u))

	fmt.Println("--- pointer interface ---")
	methods(reflect.TypeOf(&u))
}

func main() {
	u := User{"Jack", 23}
	v := reflect.ValueOf(u)
	fmt.Println(v.FieldByName("Username").Interface())
	//2、除具体的 Int、String 等转换方法，还可返回 interface{}。只是私有非导出字段无法使用，需用 CanInterface 判断一下。
	//fmt.Println(v.FieldByName("age").Interface())
}
