package main

import (
	"fmt"
	"project2/day4interface"
	"time"
)

/**
1、接口可值可指针（两个，一个实现类，一个指针，自带指针），内部含有了一个指针，无必要将接口写成指针,通常是用指针
2、可以实现类似的多态
*/
func main() {
	//1 值类型
	var r day4interface.IRetriver
	r = day4interface.Retriver{UserAgent: "fafd"}
	fmt.Println(r)
	var imock day4interface.IRetriver
	//2 指针类型，通常用这种
	imock = &myMockRetriver{"mozile", time.Minute}
	//验证接口内部是含有指针的，有两个值
	fmt.Printf("%T %v\n", imock, imock)
	//download(imock)

	//判断具体什么对象类型，值类型还是指针类型
	inspect(r)

	//验证imock是不是实现了IRetriver
	if myMock,ok := imock.(day4interface.IRetriver); ok {
		fmt.Println("Contents:",download(myMock))
	} else {
		fmt.Println("not a mock retriver")
	}
}

func download(r day4interface.IRetriver) string {
	return r.Get("http://wwww.baidu.com")
}

type myMockRetriver struct {
	UserAgent string
	TimeOut   time.Duration
}

func (m *myMockRetriver) Get(url string) string {
	//response, err := http.Get(url)
	//if err != nil {
	//	panic(err)
	//}
	//result, err := httputil.DumpResponse(response, true)
	//defer response.Body.Close()
	//if err != nil {
	//	panic(err)
	//}
	//return string(result)
	return url
}

func inspect(r day4interface.IRetriver ) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case day4interface.Retriver:
		fmt.Println("Contents:",v.UserAgent)
	case *myMockRetriver:
		fmt.Println("Contents:",v.UserAgent)
	}
}
