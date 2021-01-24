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

func main1() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	SHandlerError(err, "ResolveUDPAddr")
	conn, err := net.ListenUDP("udp", addr)
	SHandlerError(err, "listen.udp")
	bytes := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(bytes)
	SHandlerError(err, "ReadFromUdp")
	fmt.Println("读到来自%v的内容：%s\n", remoteAddr, string(bytes[:n]))
	conn.WriteToUDP([]byte("已阅"), remoteAddr)
}

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8888")
	SHandlerError(err, "ResolveUDPAddr")
	conn, err := net.ListenUDP("udp", addr)
	SHandlerError(err, "listen.udp")
	bytes := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(bytes)
		SHandlerError(err, "ReadFromUdp")
		serverMessage := string(bytes[:n])
		fmt.Println("读到来自%v的内容：%s\n", remoteAddr, serverMessage)
		if serverMessage == "im off" {
			conn.WriteToUDP([]byte("bye bye"), remoteAddr)
		} else {
			conn.WriteToUDP([]byte("已阅"), remoteAddr)
		}
	}
}
