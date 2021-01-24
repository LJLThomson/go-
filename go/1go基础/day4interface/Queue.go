package day4interface

//存储任意类型
type Queue []interface{}

//传入任意类型参数
func (queu *Queue) Push(value interface{}) {
	//可以在方法中进行限定
	//*queu = append(*queu, value.(int))
	*queu = append(*queu, value)
}

//返回任意类型参数
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
