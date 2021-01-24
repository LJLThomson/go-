package main

import (
	"fmt"
	"project2/day9reflection"
	"reflect"
)

func main() {
	person := day9reflection.Person{"张三", 25}
	famliy := day9reflection.Family{"历史", 26, &person}
	//famliy.SetName()	调用时自动识别是普通类型还是指针
	pType := reflect.TypeOf(&famliy)
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
	fmt.Println(oType.NumMethod()) // 1 找不到Family对象为指针的情况
	fmt.Println(pType.NumMethod()) // 3 *Family Family全都可以获取

	fmt.Println(pType.Method(0).Name)
	fmt.Println(pType.Method(0).Type)
	for j := 0; j < pType.NumMethod(); j++ {
		//调用方法时，可以是普通对象也可以是指针
		method := pType.Method(j)
		//方法名称func (f *Family) toString()  string 和类型
		fmt.Println(method.Name)
		fmt.Println(method.Type)
		//方法参数,第一个Family为第一个参数
		fmt.Println(method.Type.In(0))
		fmt.Println(method.Type.NumIn())
	}
	//第i对象中第0个属性名称,属性一律对象.属性
	aName := oType.FieldByIndex([]int{2, 0}).Name
	fmt.Println(aName)
	name, b := pType.MethodByName("SetName")
	fmt.Println(name, name.Name, b)
}

func main2() {
	person := day9reflection.Person{"张三", 25}
	famliy := day9reflection.Family{"历史", 26, &person}
	pValue := reflect.ValueOf(&famliy)
	oValue := pValue.Elem()
	fmt.Println("打印对象类型名称")
	fmt.Println(oValue.Kind())
	fmt.Println("打印属性值")
	oValue.NumField()
	for i := 0; i < oValue.NumField(); i++ {
		field := oValue.Field(i)
		//正射，field是反射值，不能用
		fmt.Println(field.Interface())
		//获取type属性,类似于type
		fmt.Println(field.Type())
	}
	fmt.Println("基本类型以外对象属性值")
	aName := oValue.FieldByIndex([]int{2, 0}).Interface()
	fmt.Println(aName)
	fmt.Println("查找属性")
	age := oValue.FieldByName("Age")
	if age.IsValid() {
		fmt.Println(age)
	}
	fmt.Println("属性替换值")
	fieldValue := oValue.Field(0)
	fieldValue.SetString("李四")
	fmt.Println(oValue)
	fmt.Println(pValue)
	fmt.Println(pValue.NumMethod())
	fmt.Println("方法:" + string(pValue.NumMethod()))

	for j := 0; j < pValue.NumMethod(); j++ {
		method := pValue.Method(j)
		if method.IsValid() {
			fmt.Println(method) // method是一个地址
		}
		//如果想要通过方法改变属性值，则原方法中必须是指针
	}
	//必须是对应调用，反射严格
	setName := pValue.MethodByName("SetName")
	if setName.IsValid() {
		fmt.Println(setName.Type())
		setName.Call([]reflect.Value{reflect.ValueOf("赵六")})
		fmt.Println(oValue)
	}
	// method(i)不好用，对象有自己的排序规则，不是从上到下的
	method := pValue.Method(1)
	method.Call([]reflect.Value{reflect.ValueOf("琪琪")})
	fmt.Println(oValue)
}
