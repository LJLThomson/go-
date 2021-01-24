package day3Object

import "fmt"

/**
组合方式，类似与继承了
 */
type MyTreeNode struct {
	node *Node
}

func (myTreeNode *MyTreeNode) PostOrder() {

}

func (myTreeNode *MyTreeNode) PrintValue()  {
	if myTreeNode == nil || myTreeNode.node == nil{
		return
	}
	fmt.Println(myTreeNode.node.Value)
}
