package main

import (
	"fmt"
	"unicode/utf8"
)

/**
1、处理中文，中文三字节
2、转码，UTF-8转成unicode
3、byte转Rune
*/
func main() {
	s := "yes我爱腾讯网"
	//转unicode
	for _, b := range []byte(s) {
		fmt.Printf("%X", b)
	}
	// ch是一个rune，会出错
	for i, ch := range s {
		fmt.Printf("(%d %c)", i, ch)
	}
	//计算4字节的rune所需空间，其中一个Y也占4字节
	fmt.Println("Rune count", utf8.RuneCountInString(s))

	//转成Rune
	bytes := []byte(s)
	for len(bytes) > 0 {
		//size是指当前bytes下标,转Rune
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c", ch)
	}
	for i, ch := range []rune(s) {
		fmt.Printf("%d %c", i, ch)
	}
}
