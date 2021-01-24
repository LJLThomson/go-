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

func info(m reflect.Method) {
	t := m.Type

	fmt.Println(m.Name)

	//NumIn参数个数，Data算第一个参数
	for i, n := 0, t.NumIn(); i < n; i++ {
		fmt.Printf(" in[%d] %v\n", i, t.In(i))
	}

	//NumOut返回值个数
	for i, n := 0, t.NumOut(); i < n; i++ {
		fmt.Printf(" out[%d] %v\n", i, t.Out(i))
	}
}

func main() {
	d := new(Data)
	t := reflect.TypeOf(d)
	test, _ := t.MethodByName("Test")
	info(test)

	sum, _ := t.MethodByName("Sum")
	info(sum)
}
