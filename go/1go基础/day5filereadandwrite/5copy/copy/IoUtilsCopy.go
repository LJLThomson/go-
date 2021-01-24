package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
缓冲区读取，常用
*/
func main() {
	src, _ := os.OpenFile("ab_.txt", os.O_RDONLY, 0666)
	dst, _ := os.OpenFile("ab_.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer func() {
		src.Close()
		dst.Close()
	}()
	//
	bufferReader := bufio.NewReader(src)
	bufferWritter := bufio.NewWriter(dst)
	bytes := make([]byte, 1024*8)
	for {
		n, err := bufferReader.Read(bytes)
		if err != nil {
			if err == io.EOF && n == 0 {
				fmt.Println("文件读取完毕")
				break
			} else {
				fmt.Println("文件读取错误")
				return
			}
		} else {
			_, err := bufferWritter.Write(bytes)
			if err != nil {
				fmt.Println("写出错误")
				return
			}
		}
	}
}
