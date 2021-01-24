package main

import "fmt"

/**
1、slice添加元素
2、创建slice
3、copy slice
4、delete 元素
5、删除头尾
*/
func main() {
	//1添加数组时如果超过cap,系统会重新分配更大的底层数组，此时就是全新的底层数组（抛弃了原始数据，如果没用，会被GC)，由于值传递关系，必须接收append的值
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("arr=%v,len(arr)=%d,cap(arr)=%d\n", arr, len(arr), cap(arr))
	s1 := arr[:6]
	//arr值也会变化，超过底层数组，系统分配扩展，s2就不受arr影响了
	s2 := append(s1, 10, 20, 30, 40, 50)
	fmt.Println(s2)
	fmt.Println(arr)
	arr[0] = 100
	//s2超过底层数据，值不受影响,新分配内存
	fmt.Println(arr)
	fmt.Println(s2)

	//2定义slice,三种方式
	var s3 []int
	for i := 0; i < 10; i++ {
		fmt.Printf("arr=%v,len(arr)=%d,cap(arr)=%d\n", s3, len(s3), cap(s3))
		s3 = append(s3, i)
	}
	s4 := []int{2, 4, 6, 8}
	fmt.Printf("arr=%v,len(arr)=%d,cap(arr)=%d\n", s4, len(s4), cap(s4))
	//len为16，cap为16
	s5 := make([]int, 16)
	printSlice(s5)
	//len为10，cap为16
	s6 := make([]int, 10, 16)
	printSlice(s6)

	//3 copy拷贝
	copy(s5, s6)
	printSlice(s5)

	//4、删除元素，跳到s3中间元素
	//s3[:3] + s3[4:], s3[4:]... s3后面所有参数
	s3 = append(s3[:3], s3[4:]...)
	printSlice(s3)

	//5、删除头尾元素
	front := s3[0]
	s3 = s3[1:]
	printSlice(s3)
	fmt.Println(front)
	tail := s3[len(s3)-1]
	s3 = s3[:len(s3)-1]
	printSlice(s3)
	fmt.Println(tail)
}

func printSlice(s []int) {
	fmt.Printf("arr=%v,len(arr)=%d,cap(arr)=%d\n", s, len(s), cap(s))
}
