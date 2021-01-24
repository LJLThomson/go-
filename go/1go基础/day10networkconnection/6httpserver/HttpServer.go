package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	//func实现了类似于接口的功能
	http.HandleFunc("/shit", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("shit你妹啊"))
	})

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("请求的方法：" + request.Method))
		writer.Write([]byte("请求的内容长度" + strconv.Itoa(int(request.ContentLength))))
		writer.Write([]byte("请求的主机=" + request.Host))
		writer.Write([]byte("请求的协议" + request.Proto))
		writer.Write([]byte("请求的远程地址=" + request.RemoteAddr))
		writer.Write([]byte("请求的路由" + request.RequestURI))
		writer.Write([]byte("已收到请求"))

	})

	http.HandleFunc("/hupu", func(writer http.ResponseWriter, request *http.Request) {
		//将html文件传给客户端
		bytes, err := ioutil.ReadFile("D:/IDEA/Project/src/project2/hupu.html")
		if err != nil {
			fmt.Println("文件读取错误")
			writer.Write([]byte("文件不存在"))
			return
		}
		writer.Write(bytes)
	})
	//0:0:0:0:8080,任意的ip地址，ip4和ip6都行
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println("无此地址")
	}
}
