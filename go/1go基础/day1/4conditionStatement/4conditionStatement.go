package main

import "fmt"
import "io/ioutil"

func main() {
	ifelseStatement()
	ifelseStatment2()
	fmt.Println(
		grade(20),
		grade(70),
		grade(100),
		grade(101)) // 错误数据
}

func ifelseStatement() {
	const filename = "abc.txt"
	//err不为空，说明有错
	constents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", constents)
	}
}

func ifelseStatment2() {
	const filename = "abc.txt"
	//err不为空，说明有错
	if constents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", constents)
	}
}
func grade(score int) string {
	g := "f"
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "c"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("wrong score: %d", score))
	}
	return g
}
