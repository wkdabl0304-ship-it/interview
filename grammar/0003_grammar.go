package main

import "fmt"

// 总结：
// 1.计算机里的数值使用补码表示，最高位为符号位，负数会比正数多一个
// 2.如果使用变量进行数值计算，那么在运行过程才计算
// 3.如果使用常量进行数值计算，那么在编译期就计算完了，所以报错是因为将 128 赋值给了 int8

func main() {
	//var a int8 = -1
	//var b int8 = -128 / a
	//fmt.Println(b)

	const a int8 = -1
	var b int8 = -128 / a
	fmt.Println(b)
}
