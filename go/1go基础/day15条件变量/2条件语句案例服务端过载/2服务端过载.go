package main

import (
	"fmt"
	"sync"
	"time"
)

/**
服务端：接入客户端数量超过最大值，发出预警，进入阻塞等待，等待其他客户端退出一个，才能接入
客户端：受到服务端过载预警，削减客户端用户数量，并通知服务端过载预警已解除
*/
func main() {
	server := NewServer(5)
	server.Serve()
	server.Run()
}

type Server struct {
	maxConnect int
	connection int
	cond       *sync.Cond
	channel    chan string
}

func NewServer(maxConnect int) *Server {
	server := new(Server)
	server.maxConnect = maxConnect
	server.channel = make(chan string)
	server.cond = sync.NewCond(&sync.Mutex{})
	return server
}

func (server *Server) Serve() {
	go func() {
		for {
			//收到服务端过载警告
			<-server.channel
			server.cond.L.Lock()
			server.connection--
			fmt.Println("connection数量：", server.connection)
			server.cond.Signal()
			fmt.Println("过载预警已解除")
			server.cond.L.Unlock()
			<-time.After(1 * time.Second)
		}
	}()
}

func (server *Server) Run() {
	for {
		server.cond.L.Lock()
		if server.connection == server.maxConnect {
			//过载预警
			server.channel <- "过载预警"
			server.cond.Wait()
		}
		fmt.Println("可以添加客户端啦")
		server.connection++
		server.cond.L.Unlock()
		<-time.After(1 * time.Second)
	}
}
