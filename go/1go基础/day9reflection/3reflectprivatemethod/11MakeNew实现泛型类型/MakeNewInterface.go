package main

import (
	"fmt"
	"reflect"
)

var (
	Int = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)

func Make(T reflect.Type, fptr interface{}) {

	// 实际创建 slice 的包装函数。
	swap := func(in []reflect.Value) []reflect.Value {

		// --- 省略算法内容 --- //
		// 返回和类型匹配的 slice 对象。
		return []reflect.Value{
			reflect.MakeSlice(
				reflect.SliceOf(T), // slice type
				int(in[0].Int()), // len
				int(in[1].Int()), // cap
			),
		}
	}

	// 传入的是函数变量指针，因为我们要将变量指向 swap 函数。
	fn := reflect.ValueOf(fptr).Elem()

	// 获取函数指针类型，生成所需 swap function value
	v := reflect.MakeFunc(fn.Type(), swap)

	// 修改函数指针实际指向，也就是 swap
	fn.Set(v)
}

func main() {
	var makeints func(int, int) []int
	var makestrings func(int, int) []string

	// 用相同算法，生成不同类型创建函数。
	Make(Int, &makeints)
	Make(String, &makestrings)

	// 按实际类型使用。
	x := makeints(5, 10)
	fmt.Printf("%#v\n", x)

	s := makestrings(3, 10)
	fmt.Printf("%#v\n", s)
}
/**
核心是提供一个 swap 函数，其中利用 reflect.MakeSlice 生成最终 slice 对象， 因此需要传入 element type、len、cap 参数。
接下来，利用 MakeFunc 函数生成 swap value，并修改函数变量指向，以达到调 用 swap 的目的。
当调用具体类型的函数变量时，实际内部调用的是 swap，相关代码会自动转换参 数列表，并将返回结果还原成具体类型返回值。
如此，在共享算法的前提下，无须用 interface{}，无须做类型转换，颇有泛型的效果。
 */