package main

import "fmt"

// 总结：defer 遇到链式调用时，会将前面全部调用执行完，只保留最后一个

type T struct{}

func (t T) f(n int) T {
	fmt.Print(n)
	return t
}

func main() {
	var t T
	defer t.f(1).f(2)
	fmt.Print(3)
}
