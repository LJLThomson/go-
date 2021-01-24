package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
1、以只读的方式打开一个文件
2、创建带缓存的读取器，逐行读取到末尾
4=readable,2=writable,1=exectuable
0666

*/
func main1() {
	file, err := os.OpenFile("abc.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("文件打开失败")
	}
	defer func() {
		file.Close()
		fmt.Println("文件关闭")
	}()
	bufferReader := bufio.NewReader(file)
	for {
		//以\n形式读取，\n一字节
		str, err := bufferReader.ReadString('\n')
		fmt.Println("我的", str, err)
		if err != nil && err != io.EOF {
			fmt.Println("文件读取失败")
			break
		}
		fmt.Println(str)
		if err == io.EOF {
			fmt.Println("文件已到末尾")
		}
		//if err == nil {
		//	fmt.Println(str)
		//} else {
		//	if err == io.EOF {
		//		fmt.Println("文件已到末尾")
		//		break
		//	} else {
		//		fmt.Println("文件读取失败")
		//	}
		//}
	}
}

/**
// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
O_RDONLY int = syscall.O_RDONLY // open the file read-only.
O_WRONLY int = syscall.O_WRONLY // open the file write-only.
O_RDWR   int = syscall.O_RDWR   // open the file read-write.
// The remaining values may be or'ed in to control behavior.
O_APPEND int = syscall.O_APPEND // append data to the file when writing.
O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.同步，避免被覆盖
O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
*/
func main() {
	//可读可写可追加
	file, err := os.OpenFile("b.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败")
	}
	defer func() {
		file.Close()
		fmt.Println("文件关闭")
	}()
	bufferWritter := bufio.NewWriter(file)
	bufferWritter.WriteString("我的")
	bufferWritter.WriteString("地盘")
	bufferWritter.WriteString("我做主")
	bufferWritter.WriteRune('d')
	bufferWritter.WriteByte(123)
	bufferWritter.Write([]byte{123,123,234}) // utf-8， 0到255,8位
	bufferWritter.Flush()
	fmt.Println("读写完毕")
}
