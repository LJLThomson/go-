package main

import (
	"fmt"
	"net"
	"os"
)

func SHandlerError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func ioWithConn(conn net.Conn) {
	bytes := make([]byte, 1024)

	for {
		n, err := conn.Read(bytes)
		SHandlerError(err, "conn.Read(bytes)")
		clientMsg := string(bytes[0:n])
		fmt.Println("received client:%v msg:%s\n", conn.RemoteAddr(), clientMsg)
		if clientMsg == "im off" {
			conn.Write([]byte("bye bye"))
			break
		}
		conn.Write([]byte("msg received" + clientMsg))
	}
	fmt.Println("client disconnected")
}

/**
生成exe可执行文件，命令 go build Server.go
*/
func main() {
	//addr, err2 := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	//if err2 != nil {
	//	fmt.Println("找不到改地址")
	//	return
	//}
	//net.ListenTCP("tcp", addr)
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	SHandlerError(err, "net.Listen")
	for {
		conn, err := listen.Accept()
		SHandlerError(err, "listener.Accept")
		go ioWithConn(conn)
	}
}
