package day6funcprogram

import "fmt"

type MyNode struct {
	Value       int
	Left, Right *MyNode
}

func (node *MyNode) Travese() {
	if node == nil {
		return
	}
	node.Left.Travese()
	node.print3()
	node.Right.Travese()
}

func (node *MyNode) print3() {
	if node == nil {
		return
	}
	fmt.Print(node.Value)
}

/**
函数接口编程
*/
func (node *MyNode) TraverseFunc(f func(*MyNode)) {
	if node == nil {
		return
	}
	node.Left.Travese()
	f(node)
	node.Right.Travese()
}

func (node *MyNode) Travese2() {
	node.TraverseFunc(func(n *MyNode) {
		n.print3()
	})
	fmt.Println()
}
