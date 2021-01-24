package day3Object


type MyNode MyTreeNode

func (mynode *MyTreeNode) SetValue(value int)   {
	mynode.PrintValue()
}
