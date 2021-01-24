package main

import (
	"fmt"
	"reflect"
)

type User struct {
}

type Admin struct {
	U User
	Name string
}

func (u *User) ToString() {}

//私有方法反射不到
func (a Admin) test() {}

func (a *Admin) GetUser ()  {

}

func (a *Admin) SetUser()  {

}
func (u User) SetUser()  {

}

func main() {
	u := Admin{}
	methods := func(t reflect.Type) {
		for i, n := 0, t.NumMethod(); i < n; i++ {
			m := t.Method(i)
			fmt.Println(m.Name)
		}
	}
	//对象只能获取对象的方法
	fmt.Println("--- value interface ---")
	methods(reflect.TypeOf(u))

	//指针可以获取所有的方法
	fmt.Println("--- pointer interface ---")
	methods(reflect.TypeOf(&u))
}

