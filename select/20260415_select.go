package main

import "fmt"

// 1.只有当通道关闭并且缓冲区没有数据的时候，ok 才为 false

func main() {
	ch := make(chan int, 3)
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	v, ok := <-ch
	fmt.Println("读取数据 v:", v, ", ok:", ok)
	close(ch)
	fmt.Println("ch 已关闭")
	v, ok = <-ch
	fmt.Println("读取数据 v:", v, ", ok:", ok)
	v, ok = <-ch
	fmt.Println("读取数据 v:", v, ", ok:", ok)

	v, ok = <-ch
	fmt.Println("读取数据 v:", v, ", ok:", ok)
}
