package main

import (
	"fmt"
	"strings"
)

/**
检索字符串
格式化
比较大小
裁剪
炸碎
拼接
*/
func main1() {
	//包含
	fmt.Println(strings.Contains("Hello", "el"))
	fmt.Println(strings.ContainsAny("Hello", "assHel")) // true 任意字符

	fmt.Printf("%U\n", 'h')
	fmt.Printf("%c\n", 0x0068)
	fmt.Println(strings.ContainsRune("Hello", 'h'))    // true
	fmt.Println(strings.ContainsRune("Hello", 0x0068)) // true,0x0068就是'h'

	//找不到为-1
	fmt.Println(strings.Index("fuckyou", "fuck"))      //0
	fmt.Println(strings.Index("fuckyou", "yo"))        //4
	fmt.Println(strings.IndexAny("fuckoff", "ashole")) // 4
	fmt.Println(strings.IndexRune("我爱你", '爱'))         //1
}

//格式化
func main2() {
	upper := strings.ToUpper("hello")
	fmt.Println(upper)
	fmt.Println(strings.ToLower("HELLO"))
	fmt.Println(strings.Title("hello")) // 首字符大写
}

//比较
func main3() {
	//1 表示大于， 0等于，-1小于
	compare := strings.Compare("Hello", "Hssfa")
	fmt.Println(compare)
}

//裁剪
func main4() {
	//去掉头尾空格
	space := strings.TrimSpace(" He llo ")
	fmt.Println(space)
	//去掉前缀
	prefix := strings.TrimPrefix("合：我的世界", "合")
	fmt.Println(prefix)
	suffix := strings.TrimSuffix("合：我的世界", "世界")
	fmt.Println(suffix)

	trim := strings.Trim("Fuckyouareokuc", "Fuck")
	fmt.Println(trim)
	fmt.Println(strings.TrimLeft("Fuckyouareokuc", "Fuck"))
	fmt.Println(strings.TrimRight("Fuckyouareokuc", "Fuck"))

	//根据filter,true去掉，false不去掉
	fmt.Println(strings.TrimFunc("fuckyou ", filter))
	strings.TrimFunc("fuckyou ", func(r rune) bool {
		if r == 'f' {
			return true
		} else {
			return false
		}
	})
}

//炸碎和拼接,split,分隔符
func main5() {
	split := strings.Split("fuck you fuck me", " ")
	for lenth, s := range split {
		fmt.Println(lenth, s)
	}

	//炸几次,-1无穷次
	splitN := strings.SplitN("fuck you fuck me", " ", -1)

	for lenth, s := range splitN {
		fmt.Println(lenth, s)
	}

	n := strings.SplitN("fuck you fuck me", " ", 1)

	for lenth, s := range n {
		fmt.Println(lenth, s)
	}

	//将分隔符也加进去
	after := strings.SplitAfter("fuck,you,fuck,me", ",")
	fmt.Println(after, len(after))
	afterN := strings.SplitAfterN("fuck,you,fuck,me", ",", 2)
	fmt.Println(after, len(afterN))
}

//拼接
func main() {
	join := strings.Join([]string{"fuck", "you", "fuck", "me"}, "*")
	fmt.Println(join)
}
func filter(char rune) bool {
	if char == 'f' {
		return true
	} else {
		return false
	}
}
