package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func CHandlerError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}
func main1() {
	//拨号请求链接
	conn, err := net.Dial("udp", "127.0.0.1:8888")
	CHandlerError(err, "net.Dial")
	defer conn.Close()
	n, err := conn.Write([]byte("只有main包下的main函数才能运行"))
	CHandlerError(err, "conn.write")
	fmt.Println("成功写入%d个字节\n", n)

	//创建缓存区
	buffer := make([]byte, 1024)
	//从连接中读取数据
	n, err = conn.Read(buffer)
	CHandlerError(err, "conn.read")
	fmt.Println(string(buffer[:n]))
}

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8888")
	CHandlerError(err, "net.Dial")
	reader := bufio.NewReader(os.Stdin)
	defer conn.Close()
	//创建缓存区
	buffer := make([]byte, 1024)
	for {
		line, _, _ := reader.ReadLine()
		conn.Write(line)
		//从连接中读取数据
		n, err := conn.Read(buffer)
		CHandlerError(err, "conn.read")
		serverMessage := string(buffer[0:n])
		fmt.Println("server", serverMessage)
		if serverMessage == "bye bye" {
			break
		}
	}
	fmt.Println("game over")
}
