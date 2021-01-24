package main

import "fmt"

/**
1、模拟java中toString()方法
*/
func main() {
	var sinterface fmt.Stringer
	sinterface = &systemInterface{Name: "wangwu"}
	fmt.Println(sinterface.String())
}

type systemInterface struct {
	Name string
	Age  int
}

func (s *systemInterface) String() string {
	//对象可以直接点 s.Name 等价于(*s).Name
	return fmt.Sprintf("systemInterface:{name = %s}", s.Name)
}
