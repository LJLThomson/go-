package day3Object

type Queue []int

func (queu *Queue) Push(value int) {
	*queu = append(*queu, value)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
