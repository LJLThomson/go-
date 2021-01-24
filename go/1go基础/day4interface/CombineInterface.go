package day4interface

type interfaceA interface {
	Get(url string) string
	Post()
}

type interfaceB interface {
	Read()
	Write()
}


/**
1、组合接口 interfaceA,interfaceB，非常灵活
2、继承InterfaceC时，可以获取到interfaceA,interfaceB的方法,
*/
type InterfaceC interface {
	interfaceA
	interfaceB
}
