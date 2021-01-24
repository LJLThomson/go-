package main

import "fmt"

/**
案例：寻找最长不含有重复字符串的字串，比如
abcda  最长为4
acda   最长为3
ababc  最长为3
*/
func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcda"))
	fmt.Println(lengthOfNonRepeatingSubStr("acda"))
	fmt.Println(lengthOfNonRepeatingSubStr("ababc"))
	//创建一个可以存储很多值的map集合
	firstOccurred := make(map[int]int)
	firstOccurred[1] = 3
	firstOccurred[0] = 2
	firstOccurred[4] = 5
	fmt.Println(firstOccurred)
}

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		if lastId, ok := lastOccurred[ch]; ok && lastId >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		//更新重复值下标
		lastOccurred[ch] = i
	}
	//fmt.Println(lastOccurred)
	return maxLength
}

func lengthOfNonRepeatingSubStr2(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastId, ok := lastOccurred[ch]; ok && lastId >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		//更新重复值下标
		lastOccurred[ch] = i
	}
	//fmt.Println(lastOccurred)
	return maxLength
}
