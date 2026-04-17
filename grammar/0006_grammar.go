package main

import "fmt"

// 总结：
// 1.代码走到 defer 的时候会有一个快照，记录下函数传入的参数值（只针对传入值）
// 2.如果 defer 的函数内部直接使用外部变量，那么因为共享内存，所以后续代码如果修改了值，defer 运行的时候也会随之改变

func f(n int) {
	defer fmt.Println(n)
	n += 100
}

func main() {
	f(1)

	n := 1
	defer func() {
		fmt.Println(n)
	}()
	n += 100
}
