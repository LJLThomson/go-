package day4interface

type MyQueue []interface{}

func (queu *MyQueue) Push(value int) {
	*queu = append(*queu, value)
}

func (q *MyQueue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	//强转为int类型
	return head.(int)
}

func (q *MyQueue) IsEmpty() bool {
	return len(*q) == 0
}
