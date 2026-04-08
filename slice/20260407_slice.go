package main

import "fmt"

// 1.切片 copy 的时候需要保证长度足够，而非容量
// 2.for range 获取的切片元素值是拷贝值

func f(s []int) []int {
	s = append(s, 100)
	return s
}

func main() {
	//slice1 := make([]int, 3, 5)
	//slice2 := make([]int, 5, 5)
	//slice1[0] = 1
	//slice1[1] = 2
	//slice1[2] = 3
	//copy(slice2, slice1)
	//fmt.Println(slice1, slice2)
	s := []int{0, 0}
	newS := f(s)
	fmt.Println(newS)
}
