package main

import "os"

func main() {
	src, _ := os.OpenFile("ab_.txt", os.O_RDONLY, 0666)
	dst, _ := os.OpenFile("ab_.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer func() {
		src.Close()
		dst.Close()
	}()
}
