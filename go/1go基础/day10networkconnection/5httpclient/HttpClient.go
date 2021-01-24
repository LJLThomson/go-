package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func CHandlerError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

/**
Get请求
*/
func mainGet() {
	url := "http://www.baidu.com/s?wd=肉丝"
	resp, err := http.Get(url)
	CHandlerError(err, "http.Get")
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	CHandlerError(err, "ioutil.ReadAll")
	fmt.Println("获取的内容：" + string(bytes))
}

/**
post请求
*/
func main() {
	//url := "https://httpbin.org/post?name=nimei&age=25"
	//可以不用get方式加数据
	url := "https://httpbin.org/post?"
	//contentType和读取数据方式，传递的数据可以是文件
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader("rmb=0.5&hobby=接客和约汉"))
	CHandlerError(err, "http.Post")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	CHandlerError(err, "ioutil.ReadAll")
	fmt.Println("获取的内容：" + string(bytes))
}
