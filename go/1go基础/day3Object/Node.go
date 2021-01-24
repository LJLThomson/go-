package day3Object

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) setValue(Value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node.Ignored")
		return
	}
	node.Value = Value
}

func (node *Node) print2() {
	fmt.Println(node.Value)
}
func (node *Node) print3() {
	if node == nil {
		return
	}
	fmt.Print(node.Value)
}

func (node *Node) travese() {
	if node == nil {
		return
	}
	node.Left.travese()
	node.print3()
	node.Right.travese()
}
