package main

import "fmt"

/**
变量定义
	1、variableZeroValue 命名
	2、variableInitValue  命名初始化
	3、variableTypeDeduction 自动识别类型，省略类型
	4、variableShorter   := 简短号命名
	5、外部var命名初始化
	总结：每个命名都有初始化值，安全
*/
func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa,bb,dd,gg)
}

func variableZeroValue() {
	var s string
	var a int
	fmt.Printf("%s %d\n", s, a)
}

func variableInitValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Printf("%d %d %s\n", a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

//定义外部变量，注意go没有全局变量说明，只在包类使用，更加安全
var aa = 3
var bb = 4
var kk = "gg"

//cc := 5 外部错误书写方式
//建议写法
var (
	dd = 3
	ff = 4
	gg = "gg"
)
