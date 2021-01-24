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

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	CHandlerError(err, "net.conn")
	reader := bufio.NewReader(os.Stdin)
	bytes := make([]byte, 1024)
	for {
		linebytes, _, _ := reader.ReadLine()
		conn.Write(linebytes)
		n, _ := conn.Read(bytes)
		serverMessage := string(bytes[0:n])
		fmt.Println("server", serverMessage)
		if serverMessage == "bye bye" {
			break
		}
	}
	fmt.Println("game over")
}
