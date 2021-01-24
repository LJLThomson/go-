package main

import (
	"fmt"
	"reflect"
)
type Data struct {
}

func (*Data) Test(x, y int) (int, int) {
	return x + 100, y + 100
}

func (*Data) Sum(s string, x ...int) string {
	c := 0
	for _, n := range x {
		c += n
	}

	return fmt.Sprintf(s, c)
}

func main1() {
	d := new(Data)
	v := reflect.ValueOf(d)

	//封装了一个方法，用于方法调用
	exec := func(name string, in []reflect.Value) {
		m := v.MethodByName(name)
		out := m.Call(in)

		for _, v := range out {
			fmt.Println(v.Interface())
		}
	}

	exec("Test", []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})

	fmt.Println("-----------------------")

	exec("Sum", []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})
}

/**
变量打包成slice类型
reflect.ValueOf([]int{1, 2}), // 将变参打包成 slice。
 */
func main() {
	d := new(Data)
	v := reflect.ValueOf(d)

	m := v.MethodByName("Sum")

	in := []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf([]int{1, 2}), // 将变参打包成 slice。
	}

	out := m.CallSlice(in)

	for _, v := range out {
		fmt.Println(v.Interface())
	}
}