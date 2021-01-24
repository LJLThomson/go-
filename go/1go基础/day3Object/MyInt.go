package day3Object

import "fmt"

type MyInt int

func (myint *MyInt) setInt(v int) {
	fmt.Println(v)
}
