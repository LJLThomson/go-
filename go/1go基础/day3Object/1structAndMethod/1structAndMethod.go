package main

import (
	"fmt"
)

/**
1、go面向对象，没有继承和多态，有封装,没有class，只有struct
2、创建对象和对象数组
3、工厂模式构建对象
4、构造set方法
5、遍历函数
*/

func main() {
	//2、创建对象
	root := treeNode{value: 3} //1

	root3 := new(treeNode)     //2
	fmt.Println(root, root3)   //3
	var root2 treeNode
	root2.value = 3
	//C++中指针root2.right->left,go中没有，全部用.
	root2.left = &treeNode{0, nil, nil}
	root2.right = &treeNode{1, nil, nil}
	//修改root.right.left中的值
	root2.right.left = new(treeNode)
	//C++中方式
	//(*(root2.right)).left = &treeNode{0, nil, nil}
	//(*(root2.right)).left = new(treeNode)
	fmt.Println(root2)

	nodes := []treeNode{
		{value: 3, left: nil, right: nil},
		{},
		root,
	}
	fmt.Println(nodes)
	//3
	createNode := createTreeNode(3)
	fmt.Println(createNode)
	//createPointerNode := createPointer(3)
	var createPointerNode *treeNode
	createPointerNode = createPointer(3)
	fmt.Println(*createPointerNode)
	//自动识别，
	createPointerNode2 := createPointer(4)
	fmt.Println(*createPointerNode2)

	//4
	createPointerNode2.setValue(5)
	fmt.Println(*createPointerNode2)
	createPointerNode2.print2()

	createNode.travese()
}

/**
1、创建struct
*/
type treeNode struct {
	value       int
	left, right *treeNode
}

/**
3
*/
func createTreeNode(value int) treeNode {
	return treeNode{value: value}
}

func createPointer(value int) *treeNode {
	return &treeNode{value: value}
}

/**
4
*/
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node.Ignored")
		return
	}
	node.value = value
}

func (node *treeNode) print2() {
	fmt.Println(node.value)
}
func (node *treeNode) print3() {
	if node == nil {
		return
	}
	fmt.Print(node.value)
}

/**
值遍历,先左后右边
 */
func (node *treeNode) travese() {
	if node == nil {
		return
	}
	node.left.travese()
	node.print3()
	node.right.travese()
}