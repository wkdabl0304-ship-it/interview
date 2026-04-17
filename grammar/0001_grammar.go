package main

import "fmt"

// 总结：
// 1.常量组中如果某一行的常量声明省略了类型和初始化，那么将采用上一行对应位置的类型和初始化
// 2.之所以能识别出来，是因为中间有逗号，但是末尾无逗号

func main() {
	const (
		a, b = "golang", 100
		d, e
		f bool = true
		g
	)
	fmt.Println(a, b, d, e, f, g)
}
