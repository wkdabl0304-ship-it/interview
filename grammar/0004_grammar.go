package main

import "fmt"

// 总结：
// 1.因为变量遮蔽，在函数内部作用域创建与外部同名的变量，会遮蔽外部变量，作用域结束后恢复
// 2.Goland 编译器能够识别出来并对遮蔽变量进行标绿

func main() {
	var err error
	if err == nil {
		err := fmt.Errorf("err")
		fmt.Println(1, err)
	}
	if err != nil {
		fmt.Println(2, err)
	}
}
