package main

import "fmt"

// 返回值是 0，因为在进入函数的时候返回值就已经默认初始化了

func divide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic("something went wrong") // 手动 panic 的使用方法
	return a / b
}

func main() {
	a, b := 5, 0
	fmt.Println(divide(a, b))
}
