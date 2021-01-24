package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(
		convertToBin(5), // 101
		convertToBin(13),
		convertToBin(0)) //1101
}

/**
转二进制
*/
func convertToBin(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	lsb := 0
	for ; n > 0; n /= 2 {
		lsb = n % 2
		//lsb转成字符串
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	//for代替while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/**
for死循环
 */
func forLoop() {
	for {
		fmt.Println("abc")
	}
}
