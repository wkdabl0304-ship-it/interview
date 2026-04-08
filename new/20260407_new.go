package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// new 在 Go 1.26 的新特性：
// 1.new(变量) 或 老方法 new(类型)
// 2.new 是先拷贝值再取地址，而 & 是直接取地址

func main() {
	p1 := new(Person{
		Name: "Jack",
		Age:  18,
	})
	fmt.Println(p1)
	n1 := new(int(200)) // 强制类型转换
	fmt.Println(*n1)
}
