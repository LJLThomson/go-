package main

import (
	"fmt"
	"project2/day9reflection"
	"reflect"
)

func main2() {
	animal := day9reflection.Animal{"小熊", "动物"}
	human := day9reflection.NewHuman("张三", 25, &animal)
	oType := reflect.TypeOf(human)
	fmt.Println(oType)
	fmt.Println(oType.Name())
	fmt.Println(oType.Kind())
	//fmt.Println(oType.NumField())
	//fmt.Println(oType.NumMethod())
}

/**
不含有指针的情况，type
*/
func main3() {
	animal := day9reflection.Animal{"小熊", "动物"}
	teacher := day9reflection.Teacher{"张三", 25, animal}
	oType := reflect.TypeOf(teacher)
	fmt.Println(oType)
	fmt.Println(oType.Name())
	fmt.Println(oType.Kind())
	fmt.Println()
	fmt.Println(oType.NumField())
	fmt.Println(oType.Field(0).Name)
	fmt.Println(oType.Field(0).Type)

	fmt.Println("方法")
	fmt.Println(oType.NumMethod())
	fmt.Println(oType.Method(0).Name)
	fmt.Println(oType.Method(0).Type)
	fmt.Println(oType.Method(1).Name)
	fmt.Println(oType.Method(1).Type)
	//有几个方法参数，
	fmt.Println(oType.Method(1).Type.NumIn())
	//参数类型， (t Teacher)也算一个参数
	fmt.Println(oType.Method(1).Type.In(1))
	fmt.Println()

	name, b := oType.FieldByName("Name")
	fmt.Println(name, name.Name, name.Type, b)

	//第i对象中第0个属性名称
	aName := oType.FieldByIndex([]int{2, 0}).Name
	fmt.Println(aName)
}

/**
value体系
*/
func main() {
	animal := day9reflection.Animal{"小熊", "动物"}
	teacher := day9reflection.Teacher{"张三", 25, animal}
	oValue := reflect.ValueOf(teacher)
	//反射
	fmt.Println(oValue.Field(0))
	//正射形式获取值
	name := oValue.Field(0).Interface()
	fmt.Println(name)
	fmt.Println()
	aName := oValue.FieldByIndex([]int{2, 0}).Interface()
	fmt.Println(aName)
	fmt.Println()
	byName := oValue.FieldByName("Name")
	//断言为string类型
	fmt.Println(oValue,byName)

	method := oValue.Method(1)
	fmt.Println(method)
	method.Call([]reflect.Value{reflect.ValueOf("李四")})
	fmt.Println(oValue)


}
