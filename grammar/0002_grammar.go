package main

import "fmt"

// 总结：
// 1.变量/常量 在初始化的时候如果没有声明类型，那么后续的使用中可以隐式转换
// 2.如果声明了类型，那必须显示转换

func main() {
	const N = 100
	var x int = N
	const M int32 = 100
	var y int = M
	fmt.Println(x, y)
}
